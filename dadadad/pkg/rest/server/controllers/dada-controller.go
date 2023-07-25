package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mahendrabagul/dadadad/dadadad/pkg/rest/server/models"
	"github.com/mahendrabagul/dadadad/dadadad/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type DadaController struct {
	dadaService *services.DadaService
}

func NewDadaController() (*DadaController, error) {
	dadaService, err := services.NewDadaService()
	if err != nil {
		return nil, err
	}
	return &DadaController{
		dadaService: dadaService,
	}, nil
}

func (dadaController *DadaController) CreateDada(context *gin.Context) {
	// validate input
	var input models.Dada
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger dada creation
	if _, err := dadaController.dadaService.CreateDada(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Dada created successfully"})
}

func (dadaController *DadaController) UpdateDada(context *gin.Context) {
	// validate input
	var input models.Dada
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger dada update
	if _, err := dadaController.dadaService.UpdateDada(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Dada updated successfully"})
}

func (dadaController *DadaController) FetchDada(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger dada fetching
	dada, err := dadaController.dadaService.GetDada(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, dada)
}

func (dadaController *DadaController) DeleteDada(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger dada deletion
	if err := dadaController.dadaService.DeleteDada(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Dada deleted successfully",
	})
}

func (dadaController *DadaController) ListDadas(context *gin.Context) {
	// trigger all dadas fetching
	dadas, err := dadaController.dadaService.ListDadas()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, dadas)
}

func (*DadaController) PatchDada(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*DadaController) OptionsDada(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*DadaController) HeadDada(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
