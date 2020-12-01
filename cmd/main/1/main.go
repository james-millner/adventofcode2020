package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("/Users/jamesmillner/Workspace/adventofcode2020/cmd/main/1/input/input.txt")
	if err != nil {
		log.Printf("Can't read file...")
	}
	lines := strings.Split(string(content), "\n")

	log.Printf("%v", lines)

	
}