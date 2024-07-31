package domain

import "errors"

var (
	ErrExoplanetNotFound  = errors.New("no exoplanet found with the specified name")
	ErrExoplanetsNotFound = errors.New("no exoplanet registered")
	ErrExoplanetNameEmpty = errors.New("exoplanet name cannot be empty")
	ErrExoplanetExists    = errors.New("already exoplanet registered with the same name")
	ErrIvalidType         = errors.New("supported exoplanets are only GasGiant or Terrestrial")
)
