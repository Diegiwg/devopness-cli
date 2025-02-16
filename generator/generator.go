package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

type SpecRequestBody struct{
	Content map[string]struct {
		Schema struct {
			Ref string `json:"$ref"`
		} `json:"schema"`
	} `json:"content"`
}

type SpecOperation struct {
	Summary     string   `json:"summary"`
	OperationId string   `json:"operationId"`
	Tags        []string `json:"tags"`
	// Parameters  []struct{}          `json:"parameters"`
	// Responses   map[string]struct{} `json:"responses"`
	RequestBody SpecRequestBody `json:"requestBody,omitempty"`

	Path   string
	Method string
	Arguments []string
}

type Spec struct {
	Paths map[string]map[string]SpecOperation `json:"paths"`

	ServiceGroups []struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	} `json:"x-tagGroups"`

	Components struct {
		Schemas map[string]interface{} `json:"schemas"`
	} `json:"components"`

	Services     map[string][]SpecOperation
	TagToService map[string]string
}

func LoadSpecFromFile(filePath string) Spec {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var spec = Spec{
		Services:     map[string][]SpecOperation{},
		TagToService: map[string]string{},
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

func (spec *Spec) parseServices() {
	spec.Services = map[string][]SpecOperation{}

	for path, methods := range spec.Paths {
		for method, _ := range methods {
			model := spec.Paths[path][method]

			model.Arguments = spec.GetArguments(model)

			model.Path = path
			model.Method = method

			service := spec.TagToService[model.Tags[0]]

			if service == "" {
				println("Unknown tag: " + model.Tags[0])
				continue
			}

			serviceName := strings.ReplaceAll(service, " ", "")

			spec.Services[serviceName] = append(spec.Services[serviceName], model)
		}
	}

	for service, operations := range spec.Services {
		sort.Slice(operations, func(i, j int) bool {
			return operations[i].OperationId < operations[j].OperationId
		})

		spec.Services[service] = operations
	}
}

func (spec *Spec) generateServices() {
	for service, operations := range spec.Services {
		data := map[string]interface{}{
			"Service":    service,
			"Operations": operations,
		}

		TemplateToFile("generator/templates/service.tmpl", "generated/"+service+".go", data)
	}
}

func (spec *Spec) GetArguments(operation SpecOperation) []string {
	var arguments []string

	if operation.RequestBody.Content == nil {
		return arguments
	}

	for _, content := range operation.RequestBody.Content {
		ref := strings.Replace(content.Schema.Ref, "#/components/schemas/", "", 1)

		schema := spec.Components.Schemas[ref]

		if schema == nil {
			println("Schema not found: " + ref)
			continue
		}

		schemaMap := schema.(map[string]interface{})

		properties := schemaMap["properties"].(map[string]interface{})

		for name, _ := range properties {
			arguments = append(arguments, name)
		}
	}

	return arguments
}


func TemplateToFile(templatePath string, filePath string, data interface{}) {
	functions := template.FuncMap{
		"capitalize": func(s string) string {
			if len(s) == 0 {
				return s
			}
			return strings.ToUpper(string(s[0])) + s[1:]
		},

		"renderArguments": func(arguments []string, separator string) string {
			var result string

			for _, argument := range arguments {
				result += "_" + argument + " string" + separator
			}

			return strings.TrimSuffix(result, separator)
		},
	}

	templateName := filepath.Base(templatePath)

	tmpl := template.New(templateName).Funcs(functions)

	tmpl, err := tmpl.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}

func main() {
	println("Devopness CLI Generator")

	spec := LoadSpecFromFile("generator/spec.json")
	spec.parseServices()

	os.RemoveAll("generated")
	os.MkdirAll("generated", os.ModePerm)

	spec.generateServices()
}
