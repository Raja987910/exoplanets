package handler

import (
	"encoding/json"
	"exoplanets/api"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"

	"exoplanets/domain"
)

type ExoHandler struct {
	exoService ExoService
}

func ExoNewHandler(service ExoService) *ExoHandler {
	return &ExoHandler{exoService: service}
}

// function to handle StoreExoplanet
func (h *ExoHandler) StoreExoplanet(w http.ResponseWriter, r *http.Request, urlParams httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		api.Error(w, r, fmt.Errorf("read request body failed: %w", err), http.StatusBadRequest)
		return
	}
	var req domain.Exoplanets
	if err := json.Unmarshal(body, &req); err != nil {
		api.Error(w, r, fmt.Errorf("unmarshal request body failed: %w", err), http.StatusBadRequest)
		return
	}

	if len(strings.TrimSpace(req.Name)) == 0 {
		api.Error(w, r, domain.ErrExoplanetNameEmpty, 0)
		return
	}

	if req.Type != domain.GasGiant && req.Type != domain.Terrestrial {
		api.Error(w, r, domain.ErrIvalidType, 0)
		return
	}

	if _, err := h.exoService.GetExoplanet(req.Name); err == nil {
		api.Error(w, r, domain.ErrExoplanetExists, 0)
		return
	}
	h.exoService.StoreExoplanet(req.Name, req)
	api.Success(w, r, nil)
}

// function to handle GetExoplanet
func (h *ExoHandler) GetExoplanet(w http.ResponseWriter, r *http.Request, urlParams httprouter.Params) {
	name := urlParams.ByName("name")
	if len(strings.TrimSpace(name)) == 0 {
		api.Error(w, r, domain.ErrExoplanetNameEmpty, 0)
		return
	}
	exoPlanet, err := h.exoService.GetExoplanet(name)
	if err != nil {
		api.Error(w, r, err, 0)
		return
	}
	api.SuccessJson(w, r, exoPlanet)
}

// function to handle UpdateExoplanet
func (h *ExoHandler) UpdateExoplanet(w http.ResponseWriter, r *http.Request, urlParams httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		api.Error(w, r, fmt.Errorf("read request body failed: %w", err), http.StatusBadRequest)
		return
	}
	var req domain.Exoplanets
	if err := json.Unmarshal(body, &req); err != nil {
		api.Error(w, r, fmt.Errorf("unmarshal request body failed: %w", err), http.StatusBadRequest)
		return
	}
	if len(strings.TrimSpace(req.Name)) == 0 {
		api.Error(w, r, domain.ErrExoplanetNameEmpty, 0)
		return
	}
	if req.Type != domain.GasGiant && req.Type != domain.Terrestrial {
		api.Error(w, r, domain.ErrIvalidType, 0)
		return
	}
	if _, err := h.exoService.GetExoplanet(req.Name); err != nil {
		api.Error(w, r, err, 0)
		return
	}
	h.exoService.StoreExoplanet(req.Name, req)
	api.Success(w, r, nil)
}

// function to handle GetAllExoplanets
func (h *ExoHandler) GetAllExoplanets(w http.ResponseWriter, r *http.Request, urlParams httprouter.Params) {
	exoPlanet, err := h.exoService.GetAllExoplanets()
	if err != nil {
		api.Error(w, r, err, http.StatusNotFound)
		return
	}
	api.SuccessJson(w, r, exoPlanet)
}

func (h *ExoHandler) DeleteExoplanet(w http.ResponseWriter, r *http.Request, urlParams httprouter.Params) {
	name := urlParams.ByName("name")
	if len(strings.TrimSpace(name)) == 0 {
		api.Error(w, r, domain.ErrExoplanetNameEmpty, 0)
		return
	}
	if err := h.exoService.DeleteExoplanet(name); err != nil {
		api.Error(w, r, err, 0)
		return
	}
	api.Success(w, r, nil)
}

// function to handle FuelEstimation
func (h *ExoHandler) FuelEstimation(w http.ResponseWriter, r *http.Request, urlParams httprouter.Params) {
	name := urlParams.ByName("name")
	if len(strings.TrimSpace(name)) == 0 {
		api.Error(w, r, domain.ErrExoplanetNameEmpty, 0)
		return
	}
	crewCapacity := urlParams.ByName("crew-capacity")
	if len(strings.TrimSpace(crewCapacity)) == 0 {
		api.Error(w, r, domain.ErrExoplanetNameEmpty, 0)
		return
	}
	c, err := strconv.Atoi(crewCapacity)
	if err != nil {
		api.Error(w, r, err, 0)
		return
	}
	exoPlanet, err := h.exoService.GetExoplanet(name)
	if err != nil {
		api.Error(w, r, err, 0)
		return
	}
	fuelEst := h.exoService.FuelEstimation(exoPlanet, c)
	api.SuccessJson(w, r, fuelEst)
}
