build:
	@cd cmd/core && go build -o ../../bin/gosupport

# run: build
# 	@cmd/core/bin/gosupport

run: build
	@cmd/core/bin/gosupport

run-src:
	@cd cmd/core && go run .

test:
	@go test -v ./...
