build:
	@cd cmd/core && go build -o ../../bin/gosupport

run: build
	@./bin/gosupport

test:
	@go test -v ./...
