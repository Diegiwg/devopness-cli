package generator

import (
	"os"
	"path/filepath"
	"text/template"
)


func TemplateToFile(templatePath string, filePath string, data interface{}) {
	functions := template.FuncMap{
		"capitalize": Capitalize,
		"camelcase": CamelCase,
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
