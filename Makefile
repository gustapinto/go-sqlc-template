build:
	go build -o ./bin/main.out ./cmd/main.go

run: build
	./bin/main.out

generate:
	go generate
