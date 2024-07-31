package api

import (
	"encoding/json"
	"errors"
	"exoplanets/domain"
	"fmt"
	"log"
	"net/http"
)

func SuccessJson(w http.ResponseWriter, r *http.Request, data interface{}) {
	jsonMsg, err := json.Marshal(data)
	if err != nil {
		Error(w, r, fmt.Errorf("serialising response failed: %w", err), 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		Success(w, r, jsonMsg)
	}
}

func Success(w http.ResponseWriter, r *http.Request, jsonMsg []byte) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if _, err := w.Write(jsonMsg); err != nil {
		log.Printf("Error writing response: %v", err)
	}

	log.Printf(
		"%s %s %s 200",
		r.Method,
		r.RequestURI,
		r.RemoteAddr,
	)
}

func Error(w http.ResponseWriter, r *http.Request, err error, code int) {
	if code == 0 {
		code = toHTTPStatusCode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	if err == nil {
		err = fmt.Errorf("nil err")
	}
	logErr := err
	errorMsgJSON, err := json.Marshal(domain.ErrorResponse{Message: err.Error()})
	if err != nil {
		log.Println(err)
	} else {
		if _, err = w.Write(errorMsgJSON); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	}

	log.Printf(
		"%s %s %s %d %s",
		r.Method,
		r.RequestURI,
		r.RemoteAddr,
		code,
		logErr.Error(),
	)
}

func toHTTPStatusCode(err error) int {
	switch {
	case errors.Is(err, domain.ErrExoplanetNotFound):
		return http.StatusNotFound
	case errors.Is(err, domain.ErrExoplanetsNotFound):
		return http.StatusNotFound
	case errors.Is(err, domain.ErrExoplanetNameEmpty):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrExoplanetExists):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrIvalidType):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
