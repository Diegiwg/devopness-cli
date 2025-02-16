package generator

import (
	"encoding/json"
	"os"
	"strings"
)

type SpecModel struct {
	Type       string               `json:"type"`
	Properties map[string]SpecModel `json:"properties,omitempty"`
	Enum       []string             `json:"enum,omitempty"`
	Required   []string             `json:"required,omitempty"`

	OneOf []struct {
		Ref string `json:"$ref"`
	} `json:"oneOf,omitempty"`

	Ref string `json:"$ref,omitempty"`
}

type Model struct {
	Name       string
	Type       string
	Properties []ModelProperty
	Values     []string
}

type ModelProperty struct {
	Name string
	Type string
	Ref  string
	// Required bool
}

func (spec *Spec) ParseModels() {
	println("Parsing models...")

	for name, model := range spec.Components.Schemas {
		spec.Models[name] = *spec.ParseModel(name, model)
	}
}

func (spec *Spec) ParseModel(name string, sModel SpecModel) *Model {
	// println("Parsing model: " + name)
	var model = &Model{
		Name: name,
	}

	if sModel.OneOf != nil {
		spec.parserOneOf(name, sModel, model)
		return model
	}

	switch sModel.Type {
	case "object":
		spec.parseObject(name, sModel, model)
	case "array":
		spec.parseArray(name, sModel, model)
	case "string":
		spec.parseString(name, sModel, model)
	default:
		println("Skipping model: " + name)
		println("    Type: " + sModel.Type)
	}

	return model
}

func (spec *Spec) parseObject(name string, sModel SpecModel, model *Model) {
	// println("Parsing object: " + name)

	for name, rawProp := range sModel.Properties {
		prop := spec.ParseObjectProperty(name, rawProp)
		model.Properties = append(model.Properties, prop)
	}
}

func (spec *Spec) ParseObjectProperty(name string, sModel SpecModel) ModelProperty {
	// println("Parsing property: " + name)

	var prop = ModelProperty{
		Name: name,
		Type: sModel.Type,
	}

	if sModel.Ref != "" {
		prop.Ref = spec.parseRef(sModel.Ref)
	}

	return prop
}

func (spec *Spec) parseArray(name string, sModel SpecModel, model *Model) {
	// println("Parsing array: " + name)

	model.Type = "array"
}

func (spec *Spec) parseString(name string, sModel SpecModel, model *Model) {
	// println("Parsing string: " + name)

	model.Type = "string"
	
	if sModel.Enum != nil {
		spec.parseEnum(name, sModel, model)
	}

}

func (spec *Spec) parseEnum(name string, sModel SpecModel, model *Model) {
	// println("Parsing enum: " + name)

	model.Type = "enum"
	for _, enum := range sModel.Enum {
		model.Values = append(model.Values, enum)
	}
}

func (spec *Spec) parserOneOf(name string, sModel SpecModel, model *Model) {
	println("Parsing oneOf: " + name)

	model.Type = "interface{}"

	for _, oneOf := range sModel.OneOf {
		println("    " + spec.parseRef(oneOf.Ref))
	}
}

func (spec *Spec) parseRef(name string) string {
	return strings.ReplaceAll(name, "#/components/schemas/", "")
}

func (spec *Spec) DumpModels(filePath string) {
	content, err := json.MarshalIndent(spec.Models, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		panic(err)
	}
}

func (spec *Spec) GenerateModels() {
	for name, model := range spec.Models {
		data := map[string]interface{}{
			"Model": model,
		}

		TemplateToFile("generator/templates/model.tmpl", "generated/models/"+name+".go", data)
	}
}
