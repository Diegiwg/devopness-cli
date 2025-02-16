package generator

import (
	"strings"
)

type ServiceOperation struct {
	Path        string
	Method      string
	Summary     string
	OperationId string
	RequestBody string
}

func (spec *Spec) ParseServices() {
	spec.Services = map[string][]ServiceOperation{}

	for path, methods := range spec.Paths {
		for method, _ := range methods {
			service := ServiceOperation{}
			service.Path = path
			service.Method = method

			model := spec.Paths[path][method]
			service.Summary = model.Summary
			service.OperationId = model.OperationId
			service.RequestBody = spec.GetArguments(model)

			serviceId := spec.TagToService[model.Tags[0]]
			if serviceId == "" {
				println("Unknown tag: " + model.Tags[0])
				continue
			}

			serviceName := strings.ReplaceAll(serviceId, " ", "")
			spec.Services[serviceName] = append(spec.Services[serviceName], service)
		}
	}
}

func (spec *Spec) GenerateServices() {
	for service, operations := range spec.Services {
		importModels := false
		for _, op := range operations {
			if op.RequestBody != "" {
				importModels = true
				break
			}
		}

		data := map[string]interface{}{
			"Service":      service,
			"Operations":   operations,
			"ImportModels": importModels,
		}

		TemplateToFile("generator/templates/service.tmpl", "generated/services/"+service+".go", data)
	}
}

func (spec *Spec) GetArguments(operation SpecOperation) string {
	if operation.RequestBody.Content == nil {
		return ""
	}

	for _, content := range operation.RequestBody.Content {
		return spec.parseRef(content.Schema.Ref)
	}

	return ""
}
