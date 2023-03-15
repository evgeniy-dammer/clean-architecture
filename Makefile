include .env

run:
	go run cmd/app/main.go

build: clean
	go mod download
	go build -o clean-architecture cmd/app/main.go

clean:
	rm -f clean-architecture

lint:
	gofumpt -w . && gci write --skip-generated -s standard,default . && golangci-lint run

alignment:
	fieldalignment -fix ./internal/delivery/http/contact

migrcreate:
	migrate create -ext sql -dir ./migrations -seq init

migrup:
	migrate -path ./migrations -database 'postgres://clean:${DB_PASSWORD}@localhost:5432/clean?sslmode=disable' up

migrdown:
	migrate -path ./migrations -database 'postgres://clean:${DB_PASSWORD}@localhost:5432/clean?sslmode=disable' down

swagger:
	swag init --parseDependency --generalInfo ./internal/delivery/http/delivery.go --output ./docs/