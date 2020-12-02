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
	passwordDataSlice := file.RetreiveInputFileAsListOfString("daytwo.txt", box)

	sledRentalComplient := 0
	tobogganCorpComplient := 0

	for i := 0; i < len(passwordDataSlice); i++ {
		passwordEntry := passwordDataSlice[i]

		policy, character, password := extractPolicyCharacterAndPassword(passwordEntry, ":")
		policyMin, policyMax := getPolicyNumberOfOccurances(policy)
		fmt.Printf("Policy: %s, Character: %s, Password: %s, Min: %d, Max: %d\n", policy, character, password, policyMin, policyMax)

		if isPasswordSledRentalComplient(password, character, policyMin, policyMax) {
			sledRentalComplient++
		}

		if isTobogganCorpComplient(password, character, policyMin, policyMax) {
			tobogganCorpComplient++
		}
	}

	fmt.Printf("Found a total of: %d SledRental passwords\n", sledRentalComplient)
	fmt.Printf("Found a total of: %d Toboggan passwords\n", tobogganCorpComplient)
}

func extractPolicyCharacterAndPassword(passwordEntry string, sep string) (string, string, string) {
	result := strings.Fields(passwordEntry)
	return strings.Trim(result[0], " "), strings.Trim(result[1], ": "), strings.Trim(result[2], " ")
}

func getPolicyNumberOfOccurances(policy string) (int, int) {
	result := strings.Split(policy, "-")
	min, _ := strconv.Atoi(result[0])
	max, _ := strconv.Atoi(result[1])
	return min, max
}

func isPasswordSledRentalComplient(password, character string, policyMin, policyMax int) bool {
	return strings.Contains(password, character) && strings.Count(password, character) >= policyMin && strings.Count(password, character) <= policyMax
}

func isTobogganCorpComplient(password, character string, policyMinIndex, policyMaxIndex int) bool {
	containsCharacter := strings.Contains(password, character)

	policyMinCharacter := string(password[policyMinIndex-1])
	policyMaxCharacter := string(password[policyMaxIndex-1])

	return containsCharacter &&
		((policyMinCharacter == character && policyMaxCharacter != character) ||
			(policyMinCharacter != character && policyMaxCharacter == character))
}
