GENTOOL := openapi-generator-cli
GIT_HOST := github.com
GIT_USER := vbelouso
GIT_REPO := tasks
GO_SOURCE_DIR := openapi

.PHONY: all

all: clean build

gen-server-npm:
	openapi-generator-cli generate \
	--input-spec api/openapi-spec/v1/openapi.yaml \
	--generator-name go-server \
	--git-host $(GIT_HOST) \
	--git-user-id $(GIT_USER) \
	--git-repo-id $(GIT_REPO) \
	--package-name=tasks \
	--model-package models \
	--output server/gen \
	--additional-properties=outputAsLibrary=true,sourceFolder=tasks,hideGenerationTimestamp=true
build/server: server/app/main.go server/gen/tasks/*.go
	mkdir -p build
	go build -o $@ $<
run-server-local:
	go run /home/vbelouso/repo/golang/tasks/server/app