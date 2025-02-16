package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/diegiwg/devopness-cli/core"
)

// User represents the authentication service
type User struct{}

// Login authenticates with the API using the provided email and password.
func (u *User) Login(ctx *core.Context, email string, password string) error {
	// Ensure context starts as unauthenticated
	ctx.Authenticated = false

	// Make a POST request to /users/login with email and password
	response, err := ctx.Client.Post("/users/login", map[string]string{
		"email":    email,
		"password": password,
	})

	if err != nil {
		return fmt.Errorf("login request failed: %w", err)
	}

	// Check for empty response
	if response == "" {
		return errors.New("empty response from login API")
	}

	// Define structure to parse JSON response
	var respData struct {
		TokenType    string `json:"token_type"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}

	// Decode JSON response
	if err := json.Unmarshal([]byte(response), &respData); err != nil {
		return fmt.Errorf("failed to parse login response: %w", err)
	}

	// Ensure required fields are present
	if respData.AccessToken == "" || respData.RefreshToken == "" {
		return errors.New("invalid response: missing authentication tokens")
	}

	// Update authentication context
	ctx.Client.Auth.TokenType = respData.TokenType
	ctx.Client.Auth.AccessToken = respData.AccessToken
	ctx.Client.Auth.RefreshToken = respData.RefreshToken
	ctx.Client.Auth.ExpiresIn = respData.ExpiresIn

	ctx.Authenticated = true

	return nil
}
