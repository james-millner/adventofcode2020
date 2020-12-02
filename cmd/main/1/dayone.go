package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gobuffalo/packr"
	"github.com/james-millner/adventofcode2020/pkg/file"
)

func main() {
	box := packr.NewBox("./input")
	sliceOfAccountsAsString := file.RetreiveInputFileAsListOfString("dayone.txt", box)
	accounts := convertSliceOfStringToMap(sliceOfAccountsAsString)
	fmt.Printf("%+v\n", accounts)

	accountValueFromTwoNumbers := 0
	accountValueFromThreeNumbers := 0

	for i := 0; i < len(accounts); i++ {
		numberOne := accounts[i]

		for a := 0; a < len(accounts); a++ {
			numberTwo := accounts[a]

			if numberOne+numberTwo == 2020 {
				accountValueFromTwoNumbers = getTotalExpensiveValue(numberOne, numberTwo)
			}

			for b := 0; b < len(accounts); b++ {
				numberThree := accounts[b]

				if numberOne+numberTwo+numberThree == 2020 {
					accountValueFromThreeNumbers = getTotalExpensiveValue(numberOne, numberTwo, numberThree)
				}
			}
		}
	}

	if accountValueFromTwoNumbers != 0 {
		log.Printf("Two accounts together: %d", accountValueFromTwoNumbers)
	}

	if accountValueFromThreeNumbers != 0 {
		log.Printf("Three accounts together: %d", accountValueFromThreeNumbers)
	}
}

func getTotalExpensiveValue(numbers ...int) int {
	sum := 1
	for _, num := range numbers {
		sum = num * sum
	}
	return sum
}

func convertSliceOfStringToMap(stringArray []string) map[int]int {
	elementMap := make(map[int]int)
	for i := 0; i < len(stringArray); i++ {
		number, _ := strconv.Atoi(stringArray[i])
		elementMap[i] = number
	}

	return elementMap
}
