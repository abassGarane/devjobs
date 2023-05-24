.PHONY: build,run
build:
	@go build -o bin/devjobs -ldflags "-s -w" ./cmd/api/main.go

run: build
	@./bin/devjobs