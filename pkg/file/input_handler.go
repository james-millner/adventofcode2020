package file

import (
	"log"
	"strings"

	"github.com/gobuffalo/packr"
)

func RetreiveInputFileAsListOfString(fileName string, box packr.Box) []string {
	content, err := box.FindString(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(content), "\n")
}
