build:
	@cd cmd/core && go build -o ../../bin/gosupport

run: build
	@cmd/core/bin/gosupport

test:
	@go test -v ./...
