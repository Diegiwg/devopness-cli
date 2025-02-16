package service

import (
	"encoding/json"

	"github.com/diegiwg/devopness-cli/core"
)

type Project struct{}

type DevopnessProject struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProjectListResponse []DevopnessProject

func (r *ProjectListResponse) Dump() {
	prettyJSON, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}

	println(string(prettyJSON))
}

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
