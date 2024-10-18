

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=v0.1" -o ./bin/mhxy_helper ./gin_cmd/main.go

.PHONY: docker_build
# docker_build
docker_build:
	docker build -f dockerfile -t mhxy_helper:v0.0.1 .

.PHONY: docker_run
# docker_run
docker_run:
	docker run --name mhxy-helper -d -p 8080:8080 mhxy_helper:v0.0.1