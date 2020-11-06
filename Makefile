include scripts/help.mk

.PHONY: clean install test build image run run-docker stop-docker
.DEFAULT_GOAL := help

project = beatport
build   = $(shell git rev-parse --short HEAD)
image   = jeanmorais/$(project):$(build)

.PHONY: clean
clean: ##@dev Remove vendor folder and coverage report.
	rm -rf vendor filtered_coverage.out

.PHONY: install
install: clean ##@dev Download dependencies via go mod.
	GO111MODULE=on go mod download
	GO111MODULE=on go mod vendor

.PHONY: test
test:##@test Run tests and coverage.
	DOCKER_BUILDKIT=1 \
	docker build --progress=plain \
		--tag $(image) \
		--target=test \
		--file=Dockerfile .

	docker create --name $(project)-$(build) $(image)
	docker cp $(project)-$(build):/filtered_coverage.out filtered_coverage.out
	docker rm -vf $(project)-$(build)

.PHONY: build
build:##@build Run the build process using Docker.
	DOCKER_BUILDKIT=1 \
	docker build --progress=plain \
		--tag $(image) \
		--target=build \
		--file=Dockerfile .

.PHONY: image
image:##@build Generate beatport api docker image.
	DOCKER_BUILDKIT=1 \
	docker build --progress=plain \
		--tag $(image) \
		--target=image \
		--file=Dockerfile .

.PHONY: run
run:##@dev Run locally.
	go run cmd/beatport/main.go

.PHONY: run-docker
run-docker:##@dev Run docker container.
	docker run --name $(project) -p 8080:8080 $(image)

.PHONY: stop-docker
stop-docker:##@dev Remove docker container.
	docker rm -f $(project)