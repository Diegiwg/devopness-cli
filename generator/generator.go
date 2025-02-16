package main

import (
	"os"

	generator "github.com/diegiwg/devopness-cli/generator/src"
)

func main() {
	println("Devopness CLI Generator")

	os.RemoveAll("generated")

	os.MkdirAll("generated", os.ModePerm)
	os.MkdirAll("generated/models", os.ModePerm)
	os.MkdirAll("generated/services", os.ModePerm)

	spec := generator.LoadSpecFromFile("generator/spec.json")

	spec.ParseModels()
	spec.DumpModels("generated/models/spec.json")
	spec.GenerateModels()

	// spec.ParseServices()
	// spec.GenerateServices()
}
