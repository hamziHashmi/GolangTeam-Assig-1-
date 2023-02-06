package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"patients.com/Asads-APIs/models"
	"patients.com/Asads-APIs/services"
)

type PatientController struct {
	PatientService services.PatientService
}

func New(patientservice services.PatientService) PatientController {
	return PatientController{
		PatientService: patientservice,
	}
}

func (uc *PatientController) CreatePatient(ctx *gin.Context) {
	var patient models.Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.PatientService.CreatePatient(&patient)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *PatientController) GetPatient(ctx *gin.Context) {
	patientid := ctx.Param("id")
	patient, err := uc.PatientService.GetPatient(&patientid)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, patient)
}
func (uc *PatientController) GetAll(ctx *gin.Context) {
	patients, err := uc.PatientService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, patients)
}

func (uc *PatientController) UpdatePatient(ctx *gin.Context) {
	var patient models.Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.PatientService.UpdatePatient(&patient)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc PatientController) DeletePatient(ctx *gin.Context) {
	patientid := ctx.Param("id")
	err := uc.PatientService.DeletePatient(&patientid)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
func (uc *PatientController) RegisterPatientRoutes(rg *gin.RouterGroup) {
	patientroute := rg.Group("/patient")
	patientroute.POST("/create", uc.CreatePatient)
	patientroute.GET("/get/:id", uc.GetPatient)
	patientroute.GET("/getall", uc.GetAll)
	patientroute.PATCH("/update", uc.UpdatePatient)
	patientroute.DELETE("/delete", uc.DeletePatient)

}
