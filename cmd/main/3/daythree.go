package main

import (
	"fmt"

	"github.com/gobuffalo/packr"
	"github.com/james-millner/adventofcode2020/pkg/file"
)

func main() {
	box := packr.NewBox("./input")
	mapData := file.RetreiveInputFileAsListOfString("daythree.txt", box)

	fmt.Printf("%v\n", mapData)

	firstTreeChallenge := checkSlope(3, 1, mapData)
	fmt.Printf("Trees Hit: %d\n", firstTreeChallenge)

	treesHit1By1 := checkSlope(1, 1, mapData)
	treesHit3By1 := checkSlope(3, 1, mapData)
	treesHit5By1 := checkSlope(5, 1, mapData)
	treesHit7By1 := checkSlope(7, 1, mapData)
	treesHit1By2 := checkSlope(1, 2, mapData)

	totalTrees := treesHit1By1 * treesHit3By1 * treesHit5By1 * treesHit7By1 * treesHit1By2
	fmt.Printf("Total Trees: %d\n", totalTrees)

}

func checkSlope(right, down int, slopeMap []string) int {
	treesHit := 0
	column := 0

	for i := 0; i < len(slopeMap); i += down {
		slopeRow := slopeMap[i]

		if column >= len(slopeRow) {
			column %= len(slopeRow)
		}

		slopePoint := slopeRow[column]
		if isPointATree(slopePoint) {
			treesHit++
		}

		column += right
	}

	return treesHit
}

func isPointATree(point byte) bool {
	return point == '#'
}
