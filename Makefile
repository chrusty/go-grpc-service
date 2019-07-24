build:
	@mkdir -p bin
	@CGO_ENABLED=false go build -o bin/service cmd/service/main.go

mocks:
	@mkdir -p pkg/interfaces/mocks internal/storage/mocks
	@counterfeiter -o pkg/interfaces/mocks/instrumenter.go pkg/interfaces/instrumenter.go Instrumenter
	@counterfeiter -o pkg/interfaces/mocks/logger.go pkg/interfaces/logger.go Logger
	@counterfeiter -o pkg/interfaces/mocks/tracer.go pkg/interfaces/tracer.go Tracer
	@counterfeiter -o internal/storage/mocks/storer.go internal/storage/storer.go Storer

proto:
	@mkdir -p pkg/api/v1/go
	@protoc \
		-I=api/v1/proto \
		-I=/usr/local/include \
		-I=${GOPATH}/src \
		--go_out=plugins=grpc:pkg/api/v1/go \
		api/v1/proto/*.proto
