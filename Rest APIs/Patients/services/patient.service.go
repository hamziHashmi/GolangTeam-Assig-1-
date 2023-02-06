package services

import "patients.com/Asads-APIs/models"

type PatientService interface {
	CreatePatient(*models.Patient) error
	GetPatient(*string) (*models.Patient, error)
	GetAll() ([]*models.Patient, error)
	UpdatePatient(*models.Patient) error
	DeletePatient(*string) error
}
