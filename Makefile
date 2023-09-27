PWD=$(shell pwd)
SERVICE=eater-svc
MIGRATION_PATH=${PWD}/src/infrastructure/migrations
PROTOS_PATH=./src/infrastructure/protos

server:
	go run main.go

path_s:
	$(PROTOS_PATH)

migration:
	migrate create -ext sql -dir pkg/database/migrations -seq $(table)

migrateup:
	migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=public" -path ./pkg/database/migrations up

migratedown:
	migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=public" -path ./pkg/database/migrations down
add-protos-submodule:
	git submodule add https://github.com/javohir01/service-protos.git ./src/infrastructure/protos

pull-protos-submodule:
	git submodule update --recursive --remote

gen-eater-proto:
	protoc \
	--go_out=./src/application/protos \
	--go_opt=paths=import \
	--go-grpc_out=./src/application/protos \
	--go-grpc_opt=paths=import \
	-I=$(PROTOS_PATH)/eater \
	$(PROTOS_PATH)/eater/*.proto

docker: bin-lunix
	docker build --rm -t eater-svc -f ${PWD}/deploy/Dockerfile .

compose-up:
	docker-compose -f ./deploy/docker-compose.yml up

compose-down:
	docker-compose -f ./deploy/docker-compose.yml down

