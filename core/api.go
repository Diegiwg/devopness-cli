package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// References: https://api-docs.devopness.com/
type API struct {
	Auth struct {
		TokenType    string `json:"token_type"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	} `json:"auth"`

	Host string `json:"host"`
}

func NewAPI() *API {
	return &API{
		Host: "diegiwg-api.devopness.com",
	}
}

func (r *API) PerformRequest(method, finalUrl string, reqBody string) (string, error) {
	println("Performing request: " + method + " " + finalUrl)

	req, err := http.NewRequest(method, finalUrl, strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	if r.Auth.AccessToken != "" {
		authHeader := fmt.Sprintf("%s %s", r.Auth.TokenType, r.Auth.AccessToken)
		req.Header.Add("Authorization", authHeader)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return string(body), fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	return string(body), nil
}

func (r *API) Get(path string, params map[string]string) (string, error) {
	finalUrl := fmt.Sprintf("https://%s%s", r.Host, path)

	var queryParams url.Values

	if params != nil {
		queryParams = make(url.Values)
		for k, v := range params {
			queryParams.Add(k, v)
		}
		finalUrl += "?" + queryParams.Encode()
	}

	return r.PerformRequest("GET", finalUrl, "")
}

func (r *API) Post(path string, body map[string]string) (string, error) {
	finalUrl := fmt.Sprintf("https://%s%s", r.Host, path)

	jsonData, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	return r.PerformRequest("POST", finalUrl, string(jsonData))
}
