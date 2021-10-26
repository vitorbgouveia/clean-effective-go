all: setup-infra up

all-dev: setup-infra dev

clean:
	docker-compose -f ./build/package/docker-compose.yml down -v

up-infra:
	docker-compose -f ./build/package/docker-compose.yml up --build -d rabbitmq mongodb

up:
	docker-compose -f ./build/package/docker-compose.yml up --build cashback-sales-service

stop:
	docker-compose -f ./build/package/docker-compose.yml stop cashback-sales-service

dev:
	docker-compose -f ./build/package/docker-compose.yml run cashback-sales-service yarn dev

log:
	docker-compose logs -f cashback-sales-service

logs:
	docker-compose logs -f

format:
	gofmt -w ./**/*

lint:
	golint ./...

# test:
# 	docker-compose run cashback-sales-consumer yarn test
