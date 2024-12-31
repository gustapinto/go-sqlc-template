build:
	go build -o ./bin/main.out ./cmd/main.go

docker/build:
	docker build .

docker/run:
	docker build -q . | xargs docker run -p=5000:5000 --env-file=./.env

sqlc/fix-permissions:
	sudo chown -R $(USER):$(USER) internal

