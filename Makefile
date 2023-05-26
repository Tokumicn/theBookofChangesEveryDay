GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_CONFIG_FILES=$(shell find internal/conf -name *.proto)
ERR_PROTO_ERRORS=$(shell find router -name *.err.proto)

.PHONY: init
# init env
init:
	go get github.com/google/wire/cmd/wire@v0.5.0
	go install github.com/codeskyblue/fswatch@latest

.PHONY: buildEnv
# initilize build env
buildEnv:
	export GOPROXY=https://goproxy.cn
	export GOPRIVATE=*.100tal.com
	export GOSUMDB="off"
	git config --global url."ssh://git@git.100tal.com/".insteadOf https://git.100tal.com/


.PHONY: initConfig
# initilize a config file
initConfig:
	mkdir -p ./configs && cp internal/conf/config.dev.yaml ./configs/config.yaml


.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_CONFIG_FILES)


.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: buildApi
# buildApi
buildApi:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./cmd/api

.PHONY: buildAdminApi
# buildApi
buildAdminApi:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./cmd/admin

.PHONY: buildConsumer
# buildConsumer
buildConsumer:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./cmd/consumer

.PHONY: buildScript
# buildScript
buildScript:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./cmd/script


.PHONY: generate
# generate config & wire_gen
generate:
	go generate ./...


.PHONY: all
# generate all
all:
	make config;
	make generate;


.PHONY: runScript
# start script
runScript:
	fswatch --config cmd/script/.fsw.yml


.PHONY: runConsumer
# start consumer
runConsumer:
	fswatch --config cmd/consumer/.fsw.yml


.PHONY: runApi
# start api server
runApi:
	fswatch --config cmd/api/.fsw.yml


.PHONY: runAdminApi
# start api server
runAdminApi:
	fswatch --config cmd/admin/.fsw.yml


.PHONY: omitempty
# clean omiempty in pb's struct
omitempty:
	sed -i "s/,omitempty//g" router/*/*api.pb.go  && sed -i "s/,omitempty//g" admin_api/*/*api.pb.go



.PHONY: errors
# generate errors code
errors:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:. \
		--go-errors_out=paths=source_relative:. \
		$(ERR_PROTO_ERRORS)

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`


gofmt:
		echo "正在使用gofmt格式化文件..."
		gofmt -s -w ${GOFILES}
		echo "格式化完成"
govet:
		echo "正在进行静态检测..."
		go vet $(VETPACKAGES)


# gorm gen code
gormgen:
		go run main.go


test:
	sh coverage.sh
