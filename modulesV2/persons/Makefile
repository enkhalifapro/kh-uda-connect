SERVICE = confluence-integration

clean:
	rm -rf ./bin

test:
	go test ./...

build: clean
	mkdir bin
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=0.0.0" -o bin/$(SERVICE) main.go

run:
	go run main.go run

build-docker:
	docker build --build-arg VERSION=0.0.0 -t $(SERVICE) .

run-docker:
	docker run -p 8092:8092 -d $(SERVICE)