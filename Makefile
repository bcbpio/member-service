.PHONY: build

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/createmember handler/createmember/main.go