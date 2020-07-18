# Makefile for the project

# Variables
proj_root = $(shell pwd)
tools=./tools
year := $(shell date +%Y)
date := $(shell date +%Y-%m-%dT%H:%M:%S)
ldflags="-X main.version=$(version) -X main.date=$(date) -X main.year=$(year)"
githash := $(shell git rev-parse --short HEAD)

# Go Environment Variables
# enable go modules
GO111MODULE=on
# disable proxy
GOPROXY=direct
# dont consult checksums db
GOSUMDB=off


# runs all the services in a single terminal using goreman (foreman)
.PHONY: run
run: tools
	${tools}/goreman -f configs/Procfile start
   

# runs generate for all directories, passes the path to ./tools
.PHONY: generate
generate:
	@PATH=${PATH}:${proj_root}/${tools} TOOLS=${proj_root}/${tools} go generate ./...

# build the tools
.PHONY: tools
tools: ${tools}/goreman ${tools}/protoc-gen-go

# build goreman
${tools}/goreman:
	@go build -o ${tools}/goreman github.com/mattn/goreman

# build proto bufs generator for go
${tools}/protoc-gen-go: go.sum
	@go build -o ${tools}/protoc-gen-go github.com/golang/protobuf/protoc-gen-go