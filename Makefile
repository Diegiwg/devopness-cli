dev: build
	go run devopness.go ${ARGS}

build:
	go run ./generator/generator.go