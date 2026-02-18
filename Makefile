help:
	# make all    Upgrade, Generate, Format, Lint, Tests
	#
	# make go     Generate code for the Go library
	# make ts     Generate code for the Typescript library
	# make py     Generate code for the Python library
	# make dart   Generate code for the Dart library
	# make doc    Generate the documentation
	#
	# make v00X   Upgrade the patch version of the dependencies
	# make v0XX   Upgrade the minor version of the dependencies
	# make fmt    Generate code and Format code
	# make fix    Modernize and lint auto-fix
	# make test   Check build and Test
	# make cov    Browse test coverage
	# make clean  Remove code-coverage.out

.PHONY: all
all: v0XX fmt fix cov

.PHONY: go
go:
	go run codegen/main.go -go

.PHONY: ts
ts:
	go run codegen/main.go -ts

.PHONY: py
py:
	go run codegen/main.go -py

.PHONY: dart
dart:
	go run codegen/main.go -dart

.PHONY: doc
doc:
	go run codegen/main.go -doc

go.mod:
	go mod init github.com/lynxai-team/emo
	go mod tidy

go.sum: go.mod
	go mod tidy

.PHONY: v00X
v00X: go.sum
	GOPROXY=direct go get -t -u=patch all
	go mod tidy

.PHONY: v0XX
v0XX: go.sum
	go get -u -t all
	go mod tidy

.PHONY: fmt
fmt: go.mod
	go generate ./...
	go run mvdan.cc/gofumpt@latest -w -extra -l .

.PHONY: test
test: code-coverage.out
	go build ./...

.PHONY: cov
cov: code-coverage.out
	go tool cover -html code-coverage.out

code-coverage.out: go.sum *.go */*.go Makefile
	go test -race -vet all -tags=emo -coverprofile=code-coverage.out ./...

.PHONY: fix
fix:
	go fix ./... || true
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest run --fix

.PHONY: clean
clean:
	rm -vf code-coverage.out
