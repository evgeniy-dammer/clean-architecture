
run:
	go run services/contact/cmd/app/main.go

build: clean
	go mod download
	go build -o clean-architecture services/contact/cmd/app/main.go

clean:
	rm -f clean-architecture