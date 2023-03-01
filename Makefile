
run:
	go run services/contact/cmd/app/main.go

build: clean
	go mod download
	go build -o clean-architecture services/contact/cmd/app/main.go

clean:
	rm -f clean-architecture

lint:
	gofumpt -w . && gci write --skip-generated -s standard,default . && fieldalignment -fix ./internal/model && golangci-lint run