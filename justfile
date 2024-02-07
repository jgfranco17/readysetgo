# Base command
default:
    @just --list

runall CMD *ARGS:
    -cd core && {{CMD}} {{ARGS}}
    -cd service && {{CMD}} {{ARGS}}

test:
    go test -v github.com/jgfranco17/readysetgo/...

build:
    @echo "Building CLI app..."
    go mod download all
    CGO_ENABLED=0 GOOS=linux go build -o ./readysetgo service/cmd/main.go
    chmod +x ./readysetgo
    @echo "Build successful!"
