package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetJobsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetJobsHandler() Received request on /api/v1/jobs")

	queryParams := r.URL.Query()
	limit := queryParams.Get("limit")
	titleFilter := queryParams.Get("title_filter")
	locationFilter := queryParams.Get("location_filter")

	fmt.Printf("limit: %s\n", limit)
	fmt.Printf("titleFilter: %s\n", titleFilter)
	fmt.Printf("locationFilter: %s\n", locationFilter)

	//TODO: enable this once it has been completed
	//jobs, err := client.FetchJobs(limit, titleFilter, locationFilter)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}

	// TODO: remove dummy data
	dummyJobs := []map[string]string{
		{"title": "Software Engineer", "company": "Tech Corp", "location": locationFilter},
		{"title": "Data Scientist", "company": "Data Inc", "location": locationFilter},
	}

	responseJobs := []map[string]string{}
	if jobTitle != "" {
		for _, job := range dummyJobs {
			if containsIgnoreCase
		}
	}
}
