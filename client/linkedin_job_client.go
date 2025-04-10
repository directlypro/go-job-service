package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"go-job-service/model"
)

const (
	linkedInApiBaseUrl = "LINKEDIN_API_BASE_URL"
	apiKeyEnvVar       = "API_KEY"
)

var httpClient = &http.Client{
	Timeout: 15 * time.Second,
}

func GetJobs(ctx context.Context, limit, titleFilter, locationFilter string) ([]model.LinkedinJobApiResponse, error) {
	log.Printf("GetJobs() jobs from external API - Title: %s, Location: %s", titleFilter, locationFilter)
	activeEndpoint := "/active-jb-7d"

	//--- Get API Key ---
	apiKey := os.Getenv(apiKeyEnvVar)
	if apiKey == "" {
		log.Printf("No API key found in environment variable %s", apiKeyEnvVar)
		return nil, fmt.Errorf("API key is not configured %s", apiKeyEnvVar)
	}

	baseURL := os.Getenv(linkedInApiBaseUrl)
	if baseURL == "" {
		log.Printf("No API base URL found in environment variable %s", linkedInApiBaseUrl)
		return nil, fmt.Errorf("API base URL is not configured %s", linkedInApiBaseUrl)
	}

	baseUrl, err := url.Parse(baseURL)
	if err != nil {
		log.Printf("Error parsing linkedin api base URL %s: %s", baseURL, err)
		return nil, fmt.Errorf("Internal server configuration error (baseURL: %s)", baseURL)
	}
	baseUrl.Path = activeEndpoint

	params := url.Values{}
	if titleFilter != "" && locationFilter != "" && limit != "" {
		params.Add("titleFilter", titleFilter)
		params.Add("locationFilter", locationFilter)
		params.Add("limit", limit)
	}

	baseUrl.RawQuery = params.Encode()
	requestUrl := baseUrl.String()
	log.Printf("GetJobs() request url: %s", requestUrl)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return nil, fmt.Errorf("Failed to create API request: %v", err)
	}

	req.Header.Set("x-rapidapi-host", baseUrl.Host)
	req.Header.Set("x-rapidapi-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Error sending HTTP request: %v", err)
		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("Error sending HTTP request: %v", err)
		}
		return nil, fmt.Errorf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorBodyBytes, _ := ioutil.ReadAll(resp.Body)
		errorBody := string(errorBodyBytes)
		log.Printf("Error sending HTTP request: %d %s", resp.StatusCode, errorBody)
		return nil, fmt.Errorf("Error sending HTTP request: %d", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, fmt.Errorf("Error reading response body from API: %v", err)
	}

	var jobs []model.LinkedinJobApiResponse
	err = json.Unmarshal(bodyBytes, &jobs)
	if err != nil {
		log.Printf("Error unmarshalling JSON response body: %v", err)
		log.Printf("Response body: %s", string(bodyBytes))
		return nil, fmt.Errorf("Error unmarshalling JSON response body: %v", err)
	}

	log.Printf("GetJobs() Successfully got %d jobs from external API", len(jobs))
	return jobs, nil
}
