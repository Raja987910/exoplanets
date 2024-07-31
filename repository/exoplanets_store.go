package repository

import (
	"exoplanets/domain"
)

type ExoReadings struct {
	exoAssociatedReadings map[string]domain.Exoplanets
}

func NewExoReadings(exoAssociatedReadings map[string]domain.Exoplanets) ExoReadings {
	return ExoReadings{exoAssociatedReadings: exoAssociatedReadings}
}

// function to StoreExoplanet
func (m *ExoReadings) StoreExoplanet(smartMeterId string, electricityReadings domain.Exoplanets) {
	m.exoAssociatedReadings[smartMeterId] = electricityReadings
}

// function to GetExoplanet
func (m *ExoReadings) GetExoplanet(smartMeterId string) (domain.Exoplanets, error) {
	v, ok := m.exoAssociatedReadings[smartMeterId]
	if !ok {
		return domain.Exoplanets{}, domain.ErrExoplanetNotFound
	}
	return v, nil
}

// function to UpdateExoplanet
func (m *ExoReadings) UpdateExoplanet(smartMeterId string, electricityReadings domain.Exoplanets) {
	m.exoAssociatedReadings[smartMeterId] = electricityReadings
}

// function to GetAllExoplanets
func (m *ExoReadings) GetAllExoplanets() ([]domain.Exoplanets, error) {
	var exoPlanets []domain.Exoplanets
	for _, v := range m.exoAssociatedReadings {
		exoPlanets = append(exoPlanets, v)
	}
	if len(exoPlanets) == 0 {
		return nil, domain.ErrExoplanetsNotFound
	}
	return exoPlanets, nil
}

// function to DeleteExoplanet
func (m *ExoReadings) DeleteExoplanet(exoName string) error {
	_, ok := m.exoAssociatedReadings[exoName]
	if !ok {
		return domain.ErrExoplanetNotFound
	}
	delete(m.exoAssociatedReadings, exoName)
	return nil
}

// function to FuelEstimation
func (m *ExoReadings) FuelEstimation(exoplanet domain.Exoplanets, crewCapacity int) interface{} {
	var g float64
	if exoplanet.Type == domain.GasGiant {
		g = 0.5 / (exoplanet.Radius * exoplanet.Radius)
	}
	if exoplanet.Type == domain.Terrestrial {
		g = exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
	}
	fuelEst := exoplanet.DisFromEarth / (g * g) * float64(crewCapacity)
	return fuelEst
}
