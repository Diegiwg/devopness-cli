package generator

import (
	"strings"
)

type ServiceOperation struct {
	Path           string
	Method         string
	Summary        string
	OperationId    string
	RequestBody    string
	RequestReturns []struct {
		Type string
		Code string
	}
	RequestReturnsTypes []string
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

			service.RequestBody = spec.GetRequestBody(model)
			service.RequestReturnsTypes, service.RequestReturns = spec.GetRequestReturns(model)

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

			for _, Type := range op.RequestReturnsTypes {
				if Type != "string" {
					importModels = true
					break
				}
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

func (spec *Spec) GetRequestBody(operation SpecOperation) string {
	if operation.RequestBody.Content == nil {
		return ""
	}

	for _, content := range operation.RequestBody.Content {
		return spec.parseRef(content.Schema.Ref)
	}

	return ""
}

func (spec *Spec) GetRequestReturns(operation SpecOperation) ([]string, []struct {
	Type string
	Code string
}) {
	returns := make(map[string][]string)

	for code, ref := range operation.Responses {
		if ref.Content == nil {
			continue
		}

		for _, content := range ref.Content {
			if code == "204" {
				continue
			}

			if content.Schema.Type == "string" {
				returns["string"] = append(returns["string"], code)
				continue
			}

			// Assign the type as "model.<ref>" (if applicable)
			Type := "model." + spec.parseRef(content.Schema.Ref)
			returns[Type] = append(returns[Type], code)
		}
	}

	if len(returns) == 0 {
		return nil, nil
	}

	typeWithCode := make([]struct {
		Type string
		Code string
	}, 0, len(returns))

	uniqueTypes := make(map[string]struct{})

	for Type, codes := range returns {
		for _, code := range codes {
			uniqueTypes[Type] = struct{}{}

			typeWithCode = append(typeWithCode, struct {
				Type string
				Code string
			}{
				Type: Type,
				Code: code,
			})
		}
	}

	onlyTypes := make([]string, 0, len(uniqueTypes))
	for typeName := range uniqueTypes {
		onlyTypes = append(onlyTypes, typeName)
	}

	return onlyTypes, typeWithCode
}
