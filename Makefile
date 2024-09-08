build:
	go build -o app -v ./...

run:
	go build -o app -v ./... && ./app

test:
	go test -v ./...