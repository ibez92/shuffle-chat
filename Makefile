test:
	go test -cover ./...

lint:
	golangci-lint run --config=.golangci.yml ./...

build:
	go build -o ./app ./cmd/shuffle/main.go

deps:
	go mod tidy && go mod vendor
