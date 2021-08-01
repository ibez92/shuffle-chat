test:
	go test ./...

lint:
	golangci-lint run --config=.golangci.yml ./...

build:
	go build -o ./app ./cmd/shuffle/main.go
