PACKAGE  = advent-of-code-2020

init:
	go get ./...
	go get -u github.com/stretchr/testify/assert

dayone:
	go build cmd/main/1/dayone.go
	dayone

default: build