# Base command
default:
    @just --list

runall CMD *ARGS:
    -cd core && {{CMD}} {{ARGS}}
    -cd service && {{CMD}} {{ARGS}}

test:
    @go clean -testcache
    @echo "Running unittest suite..."
    go test -v github.com/jgfranco17/readysetgo/... -cover
    @echo "Ran all unit tests!"

build:
    @echo "Building CLI app..."
    go mod download all
    CGO_ENABLED=0 GOOS=linux go build -o ./readysetgo service/cmd/main.go
    chmod +x ./readysetgo
    @echo "Build successful!"

tidy:
    just runall go mod tidy
