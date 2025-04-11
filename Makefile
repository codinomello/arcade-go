run:
	cd cmd && go run main.go

build:
	cd cmd && go build -o ../bin main.go

test:
	cd cmd && go test -v ./...

default: run