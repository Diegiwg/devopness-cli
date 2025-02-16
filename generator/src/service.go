package generator

import (
	"sort"
	"strings"
)

func (spec *Spec) ParseServices() {
	spec.Services = map[string][]SpecOperation{}

	for path, methods := range spec.Paths {
		for method, _ := range methods {
			model := spec.Paths[path][method]

			model.Path = path
			model.Method = method

			// model.Arguments = spec.GetArguments(model)

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

func (spec *Spec) GenerateServices() {
	for service, operations := range spec.Services {
		data := map[string]interface{}{
			"Service":    service,
			"Operations": operations,
		}

		TemplateToFile("generator/templates/service.tmpl", "generated/services/"+service+".go", data)
	}
}

// func (spec *Spec) GetArguments(operation SpecOperation) map[string]string {
// 	var arguments = make(map[string]string)

// 	if operation.RequestBody.Content == nil {
// 		return arguments
// 	}

// 	println(operation.Path)

// 	for _, content := range operation.RequestBody.Content {
// 		ref := strings.Replace(content.Schema.Ref, "#/components/schemas/", "", 1)

// 		schema := spec.Components.Schemas[ref]

// 		if schema == nil {
// 			println("Schema not found: " + ref)
// 			continue
// 		}

// 		println(ref)

// 		schemaMap := schema.(map[string]interface{})

// 		properties := schemaMap["properties"].(map[string]interface{})

// 		for name, props := range properties {
// 			println(name)
// 			name = CamelCase(name)

// 			// Check for $ref to another schema
// 			if props.(map[string]interface{})["$ref"] != nil {
// 				println("The $ref is not supported yet -> " + ref)
// 				arguments[name] = "interface{}"
// 				continue
// 			}

// 			// Check for oneOf
// 			if props.(map[string]interface{})["oneOf"] != nil {
// 				println("The oneOf is not supported yet -> " + ref)
// 				arguments[name] = "interface{}"
// 				continue
// 			}

// 			_type := props.(map[string]interface{})["type"].(string)

// 			var mapTypes = map[string]string{
// 				"string":  "string",
// 				"number":  "float64",
// 				"integer": "int64",
// 				"boolean": "bool",
// 			}

// 			if value, ok := mapTypes[_type]; ok {
// 				arguments[name] = value
// 			} else {
// 				println("Unknown type: " + _type)
// 				os.Exit(1)
// 			}
// 		}
// 	}

// 	return arguments
// }
