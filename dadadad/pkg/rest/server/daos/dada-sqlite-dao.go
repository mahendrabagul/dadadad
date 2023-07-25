package daos

import (
	"database/sql"
	"errors"
	"github.com/mahendrabagul/dadadad/dadadad/pkg/rest/server/daos/clients/sqls"
	"github.com/mahendrabagul/dadadad/dadadad/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type DadaDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateDadas(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS dadas(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Dsadd TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewDadaDao() (*DadaDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateDadas(sqlClient)
	if err != nil {
		return nil, err
	}
	return &DadaDao{
		sqlClient,
	}, nil
}

func (dadaDao *DadaDao) CreateDada(m *models.Dada) (*models.Dada, error) {
	insertQuery := "INSERT INTO dadas(Dsadd)values(?)"
	res, err := dadaDao.sqlClient.DB.Exec(insertQuery, m.Dsadd)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("dada created")
	return m, nil
}

func (dadaDao *DadaDao) UpdateDada(id int64, m *models.Dada) (*models.Dada, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	dada, err := dadaDao.GetDada(id)
	if err != nil {
		return nil, err
	}
	if dada == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE dadas SET Dsadd = ? WHERE Id = ?"
	res, err := dadaDao.sqlClient.DB.Exec(updateQuery, m.Dsadd, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("dada updated")
	return m, nil
}

func (dadaDao *DadaDao) DeleteDada(id int64) error {
	deleteQuery := "DELETE FROM dadas WHERE Id = ?"
	res, err := dadaDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("dada deleted")
	return nil
}

func (dadaDao *DadaDao) ListDadas() ([]*models.Dada, error) {
	selectQuery := "SELECT * FROM dadas"
	rows, err := dadaDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var dadas []*models.Dada
	for rows.Next() {
		m := models.Dada{}
		if err = rows.Scan(&m.Id, &m.Dsadd); err != nil {
			return nil, err
		}
		dadas = append(dadas, &m)
	}
	if dadas == nil {
		dadas = []*models.Dada{}
	}

	log.Debugf("dada listed")
	return dadas, nil
}

func (dadaDao *DadaDao) GetDada(id int64) (*models.Dada, error) {
	selectQuery := "SELECT * FROM dadas WHERE Id = ?"
	row := dadaDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Dada{}
	if err := row.Scan(&m.Id, &m.Dsadd); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("dada retrieved")
	return &m, nil
}
