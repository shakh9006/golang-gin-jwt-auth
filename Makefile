.PHONY:build
build:
	go build -o main ./cmd/apiserver/main.go

docker-up:
	docker-compose up -d --build

docker-down:
	docker-compose down

.PHONY_DEFAULT := build