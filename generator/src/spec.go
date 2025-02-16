package generator

import (
	"encoding/json"
	"os"
)

type SpecInnerRef struct {
	Content map[string]struct {
		Schema struct {
			Type string `json:"type"`
			Ref  string `json:"$ref"`
		} `json:"schema"`
	} `json:"content"`
}

type SpecOperation struct {
	Summary     string                  `json:"summary"`
	OperationId string                  `json:"operationId"`
	Tags        []string                `json:"tags"`
	Responses   map[string]SpecInnerRef `json:"responses"`
	RequestBody SpecInnerRef            `json:"requestBody,omitempty"`
	// Parameters  []struct{}          `json:"parameters"`
}

type Spec struct {
	Paths map[string]map[string]SpecOperation `json:"paths"`

	ServiceGroups []struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	} `json:"x-tagGroups"`

	Components struct {
		Schemas map[string]SpecModel `json:"schemas"`
	} `json:"components"`

	Services     map[string][]ServiceOperation
	TagToService map[string]string
	Models       map[string]Model
}

func LoadSpecFromFile(filePath string) Spec {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var spec = Spec{
		Services:     map[string][]ServiceOperation{},
		TagToService: map[string]string{},
		Models:       map[string]Model{},
	}

	err = json.Unmarshal(content, &spec)
	if err != nil {
		panic(err)
	}

	// Parse Service Groups to TagToService
	for _, serviceGroup := range spec.ServiceGroups {
		for _, tag := range serviceGroup.Tags {
			spec.TagToService[tag] = serviceGroup.Name
		}
	}

	return spec
}
