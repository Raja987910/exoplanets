package router

import (
	"exoplanets/api"

	"exoplanets/repository"
	"fmt"
	"log"
	"net/http"
	"os"

	"exoplanets/handler"

	"github.com/julienschmidt/httprouter"
)

func NewServer() *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	return &http.Server{Addr: "localhost:" + port, Handler: newHandler()}
}

func addRoutes(r *httprouter.Router) {

	exoplanetReadings := repository.NewExoReadings(
		exoPlanetsCache(),
	)

	exoplanetsHandler := handler.ExoNewHandler(&exoplanetReadings)

	r.POST("/exoplanets/store", exoplanetsHandler.StoreExoplanet)
	r.GET("/exoplanets/read/:name", exoplanetsHandler.GetExoplanet)
	r.PUT("/exoplanets/update", exoplanetsHandler.UpdateExoplanet)
	r.GET("/exoplanets/read-all", exoplanetsHandler.GetAllExoplanets)
	r.DELETE("/exoplanets/delete/:name", exoplanetsHandler.DeleteExoplanet)
	r.GET("/exoplanets/fuel-estimation/:name/:crew-capacity", exoplanetsHandler.FuelEstimation)

}

func newHandler() http.Handler {
	r := httprouter.New()
	addRoutes(r)

	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			header := w.Header()
			header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		}
		w.WriteHeader(http.StatusNoContent)
	})

	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		log.Printf("panic: %+v", err)
		api.Error(w, r, fmt.Errorf("whoops! My handler has run into a panic"), http.StatusInternalServerError)
	}
	r.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.Error(w, r, fmt.Errorf("we have OPTIONS for youm but %v is not among them", r.Method), http.StatusMethodNotAllowed)
	})

	return r
}
