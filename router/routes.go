package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type service struct {
	httpServer http.Server
	url        string
	timeout    time.Duration
	router     *mux.Router
	jobsApi    *jobsApi
}
