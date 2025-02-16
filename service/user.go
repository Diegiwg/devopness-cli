package service

import (
	"encoding/json"

	"github.com/diegiwg/devopness-cli/core"
)

type User struct{}

func (r *User) Login(ctx *core.Context, email string, password string) error {
	ctx.Authenticated = false

	response, err := ctx.Client.Post("/users/login", map[string]string{
		"email":    email,
		"password": password,
	})

	if err != nil {
		print(response)
		return err
	}

	type loginResponse struct {
		TokenType    string `json:"token_type"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}

	var login loginResponse
	err = json.Unmarshal([]byte(response), &login)
	if err != nil {
		return err
	}

	ctx.Client.Auth.TokenType = login.TokenType
	ctx.Client.Auth.AccessToken = login.AccessToken
	ctx.Client.Auth.RefreshToken = login.RefreshToken
	ctx.Client.Auth.ExpiresIn = login.ExpiresIn

	ctx.Authenticated = true

	return nil
}
