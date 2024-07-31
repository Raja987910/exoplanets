package domain

type Exoplanets struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	DisFromEarth float64 `json:"distance_from_earth"`
	Radius       float64 `json:"radius"`
	Mass         float64 `json:"mass"`
	Type         string  `json:"type"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
