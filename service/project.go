package service

import (
	"encoding/json"

	"github.com/diegiwg/devopness-cli/core"
)

type Project struct{}

type ProjectListResponse []struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (r *ProjectListResponse) Dump() {
	prettyJSON, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}

	println(string(prettyJSON))
}

// List retrieves a list of all projects associated with the authenticated user.
func (p *Project) List(ctx *core.Context) ProjectListResponse {
	rawResponse, err := ctx.Client.Get("/projects", nil)
	if err != nil {
		panic(err)
	}

	var response ProjectListResponse
	err = json.Unmarshal([]byte(rawResponse), &response)
	if err != nil {
		panic(err)
	}

	return response
}
