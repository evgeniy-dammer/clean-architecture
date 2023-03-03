
run:
	go run cmd/app/main.go

build: clean
	go mod download
	go build -o clean-architecture cmd/app/main.go

clean:
	rm -f clean-architecture

lint:
	gofumpt -w . && gci write --skip-generated -s standard,default . &&  golangci-lint run

alignment:
	fieldalignment -fix ./internal/delivery/grpc

