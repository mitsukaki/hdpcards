build:
	CGO_ENABLED=0 go build -o bin/hdpc src/*.go

init:
	go get ./...

clean:
	rm -rf hdpc