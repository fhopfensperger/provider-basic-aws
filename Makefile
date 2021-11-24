# Set the shell to bash always
SHELL := /bin/bash

# Options
REGISTRY=quay.io
ORG_NAME=fhopfensperger
PROVIDER_NAME=provider-basic-aws
VERSION=v0.1.0

build: generate test
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ./bin/$(PROVIDER_NAME)-controller cmd/provider/main.go

image: generate test
	docker build . -t $(REGISTRY)/$(ORG_NAME)/$(PROVIDER_NAME)-controller:$(VERSION) -f cluster/Dockerfile.local

image-push:
	docker push $(REGISTRY)/$(ORG_NAME)/$(PROVIDER_NAME)-controller:$(VERSION)

image-multi-platform: generate test
	 docker buildx build \
	  --platform linux/arm64,linux/amd64 \
	  --tag $(REGISTRY)/$(ORG_NAME)/$(PROVIDER_NAME)-controller:$(VERSION) \
	  --tag $(REGISTRY)/$(ORG_NAME)/$(PROVIDER_NAME)-controller:latest \
	  --file cluster/Dockerfile.local \
	  --push .

run: generate
	kubectl apply -f package/crds/ -R
	go run cmd/provider/main.go -d

build-provider:
	cd package; kubectl crossplane build provider

push-provider:
	cd package; kubectl crossplane push provider $(REGISTRY)/$(ORG_NAME)/$(PROVIDER_NAME):$(VERSION); rm *.xpkg

all: image image-push build-provider push-provider
all-multi-platform: image-multi-platform build-provider push-provider

generate:
	go generate ./...
	@find package/crds -name *.yaml -exec sed -i.sed -e '1,2d' {} \;
	@find package/crds -name *.yaml.sed -delete

lint:
	$(LINT) run

tidy:
	go mod tidy

test:
	go test -v ./...

clean:
	go clean --modcache
	@find package/crds -name *.yaml -delete

# Tools

KIND=$(shell which kind)
LINT=$(shell which golangci-lint)

.PHONY: generate tidy lint clean build image all run
