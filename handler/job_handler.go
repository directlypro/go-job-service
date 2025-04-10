package handler

import (
	"encoding/json"
	"fmt"
	"go-job-service/client"
	"log"
	"net/http"
)

func GetJobsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetJobsHandler() Received request on /api/v1/jobs")

	queryParams := r.URL.Query()
	limit := queryParams.Get("limit")
	titleFilter := queryParams.Get("title")
	locationFilter := queryParams.Get("location")

	// we don't really have to print these.
	fmt.Printf("limit: %s\n", limit)
	fmt.Printf("title: %s\n", titleFilter)
	fmt.Printf("location: %s\n", locationFilter)

	//TODO: enable this once it has been completed
	jobs, err := client.GetJobs(r.Context(), limit, titleFilter, locationFilter)
	if err != nil {
		log.Printf("GetJobsHandler() Error getting jobs: %s\n", err)
		http.Error(w, "Failed to get job listing from teh external source", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(jobs)
	if err != nil {
		log.Printf("GetJobsHandler() Error encoding response: %v\n", err)
		http.Error(w, "Failed to get job listing from teh external source", http.StatusInternalServerError)
	}

}
