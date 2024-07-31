package handler

import (
	"exoplanets/domain"
	"exoplanets/repository"
)

type ExoService interface {
	StoreExoplanet(exoName string, reading domain.Exoplanets)
	GetExoplanet(exoName string) (domain.Exoplanets, error)
	UpdateExoplanet(exoName string, reading domain.Exoplanets)
	GetAllExoplanets() ([]domain.Exoplanets, error)
	DeleteExoplanet(exoName string) error
	FuelEstimation(explanet domain.Exoplanets, crewCpacity int) interface{}
}

type exoService struct {
	meterReadings *repository.ExoReadings
}

func (s *exoService) StoreExoplanet(exoName string, reading domain.Exoplanets) {
	s.meterReadings.StoreExoplanet(exoName, reading)
}

func (s *exoService) GetExoplanet(exoName string) (domain.Exoplanets, error) {
	return s.meterReadings.GetExoplanet(exoName)
}

func (s *exoService) UpdateExoplanet(exoName string, reading domain.Exoplanets) {
	s.meterReadings.StoreExoplanet(exoName, reading)
}

func (s *exoService) GetAllExoplanets() ([]domain.Exoplanets, error) {
	return s.meterReadings.GetAllExoplanets()
}

func (s *exoService) DeleteExoplanet(exoName string) error {
	return s.meterReadings.DeleteExoplanet(exoName)
}

func (s *exoService) FuelEstimation(exoplanet domain.Exoplanets, crewCpacity int) interface{} {
	return s.meterReadings.FuelEstimation(exoplanet, crewCpacity)
}
