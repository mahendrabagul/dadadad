package services

import (
	"github.com/mahendrabagul/dadadad/dadadad/pkg/rest/server/daos"
	"github.com/mahendrabagul/dadadad/dadadad/pkg/rest/server/models"
)

type DadaService struct {
	dadaDao *daos.DadaDao
}

func NewDadaService() (*DadaService, error) {
	dadaDao, err := daos.NewDadaDao()
	if err != nil {
		return nil, err
	}
	return &DadaService{
		dadaDao: dadaDao,
	}, nil
}

func (dadaService *DadaService) CreateDada(dada *models.Dada) (*models.Dada, error) {
	return dadaService.dadaDao.CreateDada(dada)
}

func (dadaService *DadaService) UpdateDada(id int64, dada *models.Dada) (*models.Dada, error) {
	return dadaService.dadaDao.UpdateDada(id, dada)
}

func (dadaService *DadaService) DeleteDada(id int64) error {
	return dadaService.dadaDao.DeleteDada(id)
}

func (dadaService *DadaService) ListDadas() ([]*models.Dada, error) {
	return dadaService.dadaDao.ListDadas()
}

func (dadaService *DadaService) GetDada(id int64) (*models.Dada, error) {
	return dadaService.dadaDao.GetDada(id)
}
