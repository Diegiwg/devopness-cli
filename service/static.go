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

// CredentialOptions retrieves the list of supported credential types and their
// configurations from the API.
//
// The method returns a CredentialOptionsResponse object which contains the
// following fields:
//
//	ProviderTypes: A list of supported provider types, each containing a
//	  "type" and a "type_human_readable" field.
//
//	SupportedProviders: A list of supported providers, each containing a
//	  "code", "name", "hint", "type", and "type_human_readable" field.
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
