SHELL := /bin/bash
CURRENT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
VERSION := $(shell git describe --tag --always --long)

test: upload-verify run-docker
	./upload-verify --verbose --local=${CURRENT_DIR} --url=http://127.0.0.1:31313/

upload-verify: $(wildcard *.go) $(wildcard **/*.go)
	go get ./...
	go build -o upload-verify -i -ldflags "-s -w -X main.version=${VERSION}"
	upx ${CURRENT_DIR}upload-verify || true

run-docker:
	docker rm -f nginx-upload-verify-test || true
	docker run --name nginx-upload-verify-test -v ${CURRENT_DIR}:/usr/share/nginx/html:ro -p 31313:80 -d nginx:stable-alpine
	wget --retry-connrefused --tries=5 -q --wait=1 --spider http://127.0.0.1:31313 || true