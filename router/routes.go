package router

import (
	"github.com/gorilla/mux"
	"go-job-service/handler"
	"log"
	"net/http"
)

func NewServiceRouter() *mux.Router {
	r := mux.NewRouter()

	//Route Register
	r.HandleFunc("/healthz", handler.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/jobs", handler.GetJobsHandler).Methods(http.MethodGet)

	log.Println("API v1 routes registered")
	return r
}
