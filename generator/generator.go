package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

type SpecOperation struct {
	Summary     string   `json:"summary"`
	OperationId string   `json:"operationId"`
	Tags        []string `json:"tags"`
	// Parameters  []struct{}          `json:"parameters"`
	// Responses   map[string]struct{} `json:"responses"`
}

type Spec struct {
	Paths    map[string]map[string]SpecOperation `json:"paths"`
	Services map[string][]SpecOperation
}

func LoadSpecFromFile(filePath string) Spec {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var spec Spec
	err = json.Unmarshal(content, &spec)
	if err != nil {
		panic(err)
	}

	return spec
}

func (spec *Spec) parseServices() {
	var tagToService map[string]string = map[string]string{
		"Actions":                "action",
		"Actions - Logs":         "action",
		"Environments - Actions": "action",
		"Projects - Actions":     "action",

		"Applications":                "application",
		"Applications - Deployments":  "application",
		"Applications - Hooks":        "application",
		"Applications - Variables":    "application",
		"Environments - Applications": "application",

		"Credentials":                "credential",
		"Credentials - Repositories": "credential",
		"Environments - Credentials": "credential",

		"Cron Jobs":                "cron-job",
		"Environments - Cron Jobs": "cron-job",

		"Daemons":                "daemon",
		"Environments - Daemons": "daemon",

		"Environments":                     "environment",
		"Projects - Environments":          "environment",
		"Projects - Archived Environments": "environment",

		"Hooks":            "hook",
		"Hooks - Requests": "hook",
		"Hook Requests":    "hook",

		"Networks":                "network",
		"Environments - Networks": "network",

		"Network Rules":                "network-rule",
		"Environments - Network Rules": "network-rule",

		"Pipelines":           "pipeline",
		"Pipelines - Actions": "pipeline",
		"Pipelines - Hooks":   "pipeline",
		"Pipelines - Steps":   "pipeline",

		"Projects": "project",

		"Resource Events": "resource-event",

		"Resource Links": "resource-link",

		"Roles":            "role",
		"Projects - Roles": "role",

		"Static Data - Application Options":              "static",
		"Static Data - Cloud Provider Service Instances": "static",
		"Static Data - Cloud Provider Services":          "static",
		"Static Data - Credential Options":               "static",
		"Static Data - Cron Job Options":                 "static",
		"Static Data - Environment Options":              "static",
		"Static Data - Network Rule Options":             "static",
		"Static Data - Permissions":                      "static",
		"Static Data - Resource Types":                   "static",
		"Static Data - Server Options":                   "static",
		"Static Data - Service Options":                  "static",
		"Static Data - User Profile Options":             "static",
		"Static Data - Virtual Host Options":             "static",
	}

	spec.Services = map[string][]SpecOperation{}

	for path, methods := range spec.Paths {
		for method, _ := range methods {
			println(path, method)

			model := spec.Paths[path][method]

			service := tagToService[model.Tags[0]]

			if service == "" {
				println("Unknown tag: " + model.Tags[0])
				continue
			}

			spec.Services[service] = append(spec.Services[service], model)
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

func main() {
	println("Devopness CLI Generator")

	spec := LoadSpecFromFile("generator/spec.json")
	spec.parseServices()

	os.RemoveAll("generated")
	os.MkdirAll("generated", os.ModePerm)

	spec.generateServices()
}

func TemplateToFile(templatePath string, filePath string, data interface{}) {
	functions := template.FuncMap{
		"capitalize": func(s string) string {
			if len(s) == 0 {
				return s
			}
			return strings.ToUpper(string(s[0])) + s[1:]
		},

		"camelCase": func(input string) string {
			words := strings.Split(input, "-")

			for i, word := range words {
				words[i] = strings.ToUpper(string(word[0])) + word[1:]
			}

			return strings.Join(words, "")
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
