package service

import (
	"encoding/json"

	"github.com/diegiwg/devopness-cli/core"
)

type Static struct{}

type CredentialOptionsResponse struct {
	ProviderTypes []struct {
		Type              string `json:"type"`
		TypeHumanReadable string `json:"type_human_readable"`
	} `json:"provider_types"`

	SupportedProviders []struct {
		Code              string `json:"code"`
		Name              string `json:"name"`
		Hint              string `json:"hint"`
		Type              string `json:"type"`
		TypeHumanReadable string `json:"type_human_readable"`
	} `json:"supported_providers"`
}

func (r *CredentialOptionsResponse) Dump() {
	prettyJSON, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}

	println(string(prettyJSON))
}

func (r *Static) CredentialOptions(ctx *core.Context) CredentialOptionsResponse {
	rawResponse, err := ctx.Client.Get("/static/credential-options", nil)
	if err != nil {
		panic(err)
	}

	var response CredentialOptionsResponse
	err = json.Unmarshal([]byte(rawResponse), &response)
	if err != nil {
		panic(err)
	}

	return response
}
