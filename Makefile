test:
	go test -cover ./...

lint-local:
	golangci-lint run --config=.golangci.yml ./...

lint-action:
	go run -mod=vendor github.com/golangci/golangci-lint/cmd/golangci-lint run

build:
	go build -o ./app ./cmd/shuffle/main.go

deps:
	go mod tidy && go mod vendor
