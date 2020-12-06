package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/james-millner/adventofcode2020/pkg/file"
)

func main() {
	box := packr.NewBox("./input")
	seatsSlice := file.RetreiveInputFileAsListOfString("dayfive.txt", box)

	fmt.Printf("%v\n", seatsSlice)

	//Part one
	//Find the highest seat ID
	highestSeatID, lowestID, allSeats := findHeighestSeat(seatsSlice)
	fmt.Printf("The highest seat ID: %d\n", highestSeatID)

	//Part two
	//Find the missing value which will be my seat.
	mySeat := findMySeat(lowestID, highestSeatID, allSeats)
	if mySeat != 0 {
		fmt.Printf("My seat ID: %d\n", mySeat)
	} else {
		fmt.Printf("No idea what seat you're on, stow away with the luggage.")
	}
}

func findHeighestSeat(seatsList []string) (int, int, map[int]string) {
	seatsMap := make(map[int]string, 0)
	lowestID, highestID := 0, 0

	//Part one
	//Find the highest seat ID
	for _, seat := range seatsList {
		rowNumber := getRowNumber(seat)
		colNumber := getColumnNumber(seat)

		id := rowNumber*8 + colNumber

		seatsMap[id] = seat

		if id > highestID {
			highestID = id
		} else if id < highestID {
			lowestID = id
		}
	}

	return highestID, lowestID, seatsMap
}

func findMySeat(lowestID, highestID int, seats map[int]string) int {
	for i := lowestID; i < highestID; i++ {
		_, ok := seats[i]

		if !ok {
			return i
		}
	}
	return 0
}

func getRowNumber(seat string) int {
	row := seat[:7]

	result := strings.ReplaceAll(strings.ReplaceAll(row, "F", "0"), "B", "1")
	return convertBinaryNumberToInt64(result)
}

func getColumnNumber(seat string) int {
	column := seat[7:]
	result := strings.ReplaceAll(strings.ReplaceAll(column, "L", "0"), "R", "1")
	return convertBinaryNumberToInt64(result)
}

func convertBinaryNumberToInt64(binaryNumber string) int {
	res, _ := strconv.ParseInt(binaryNumber, 2, 64)
	return int(res)
}
