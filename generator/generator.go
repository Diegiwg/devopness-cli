package main

import (
	"os"

	generator "github.com/diegiwg/devopness-cli/generator/src"
)

func main() {
	println("Devopness CLI Generator")

	spec := generator.LoadSpecFromFile("generator/spec.json")
	spec.ParseServices()

	os.RemoveAll("generated")

	os.MkdirAll("generated", os.ModePerm)
	os.MkdirAll("generated/models", os.ModePerm)
	os.MkdirAll("generated/services", os.ModePerm)

	spec.GenerateServices()
}

