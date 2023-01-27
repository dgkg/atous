#!/bin/bash

.PHONY: build
build:
	docker buildx build -t atous/uber:latest --platform linux/amd64 -f Dockerfile .
