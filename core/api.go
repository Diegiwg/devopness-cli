package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// References: https://api-docs.devopness.com/
type API struct {
	Auth struct {
		TokenType    string `json:"token_type"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int64    `json:"expires_in"`
	} `json:"auth"`

	Host string `json:"host"`
}

// NewAPI initializes a new API client with the default host.
func NewAPI() *API {
	return &API{
		Host: "diegiwg-api.devopness.com",
	}
}

// Sends an HTTP request with the specified method, URL, and body.
func (r *API) performRequest(method, finalURL string, reqBody string) (int, string) {
	fmt.Printf("Performing request: %s %s\n", method, finalURL)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, finalURL, strings.NewReader(reqBody))
	if err != nil {
		panic(err)
	}

	if r.Auth.AccessToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("%s %s", r.Auth.TokenType, r.Auth.AccessToken))
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return resp.StatusCode, string(body)
}

// Performs a DELETE request to remove a resource.
func (r *API) Delete(path string, params interface{}) (int, string) {
	finalURL := fmt.Sprintf("https://%s%s", r.Host, path)

	return r.performRequest(http.MethodDelete, finalURL, "")
}

// Performs a GET request with optional query parameters.
func (r *API) Get(path string, params interface{}) (int, string) {
	finalURL := fmt.Sprintf("https://%s%s", r.Host, path)

	return r.performRequest(http.MethodGet, finalURL, "")
}

// Performs a POST request with a JSON body.
func (r *API) Post(path string, body interface{}) (int, string) {
	finalURL := fmt.Sprintf("https://%s%s", r.Host, path)

	jsonData, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	return r.performRequest(http.MethodPost, finalURL, string(jsonData))
}

// Performs a PUT request to update a resource.
func (r *API) Put(path string, body interface{}) (int, string) {
	finalURL := fmt.Sprintf("https://%s%s", r.Host, path)

	jsonData, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	return r.performRequest(http.MethodPut, finalURL, string(jsonData))
}
