PACKAGE  = advent-of-code-2020

init:
	go get ./...
	go get -u github.com/stretchr/testify/assert

dayone:
	go build cmd/main/1/dayone.go
	dayone
	rm dayone

daytwo:
	go build cmd/main/2/daytwo.go
	daytwo
	rm daytwo

default: build