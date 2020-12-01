package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("./input")
	accounts := retreiveListOfAccountsFromFile("dayone.txt", box)

	fmt.Printf("%+v\n", accounts)

	accountValueFromTwoNumbers := 0
	accountValueFromThreeNumbers := 0

	for i := 0; i < len(accounts); i++ {
		numberOne := accounts[i]

		for a := 0; a < len(accounts); a++ {
			numberTwo := accounts[a]

			if numberOne+numberTwo == 2020 {
				accountValueFromTwoNumbers = getTotalAccountValue(numberOne, numberTwo)
			}

			for b := 0; b < len(accounts); b++ {
				numberThree := accounts[b]

				if numberOne+numberTwo+numberThree == 2020 {
					accountValueFromThreeNumbers = getTotalAccountValue(numberOne, numberTwo, numberThree)
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

func getTotalAccountValue(numbers ...int) int {
	sum := 1
	for _, num := range numbers {
		sum = num * sum
	}
	return sum
}

func retreiveListOfAccountsFromFile(fileName string, box packr.Box) map[int]int {
	content, err := box.FindString("dayone.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return convertSliceOfStringToMap(lines)
}

func convertSliceOfStringToMap(stringArray []string) map[int]int {
	elementMap := make(map[int]int)
	for i := 0; i < len(stringArray); i++ {
		number, _ := strconv.Atoi(stringArray[i])
		elementMap[i] = number
	}

	return elementMap
}
