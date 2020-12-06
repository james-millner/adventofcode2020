package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/james-millner/adventofcode2020/pkg/file"
	"github.com/james-millner/adventofcode2020/pkg/function"
)

func main() {
	box := packr.NewBox("./input")
	input := file.RetreiveInputFileAsString("daysix.txt", box)

	regex := regexp.MustCompile(`(?m)^\s*$`).ReplaceAllString(strings.TrimSpace(input), "-")
	passports := strings.Split(regex, "-")

	questionGroup := function.Map(passports, createGroupAnswer)

	totalYes := 0
	for _, v := range questionGroup {
		totalYes += countQuestionsByAllYes(v)
	}

	fmt.Printf("Total yes: %d\n", totalYes)

	totalYesForDistinctQuestions := 0
	for _, v := range questionGroup {
		totalYesForDistinctQuestions += countQuestionsWhereAllAnsweredYes(v)
	}

	fmt.Printf("Total where all people in the group said yes: %d\n", totalYesForDistinctQuestions)

}

func createGroupAnswer(group string) string {
	pieces := strings.FieldsFunc(
		group,
		func(r rune) bool { return r == '\n' || r == ' ' },
	)

	return strings.Join(pieces, ",")
}

func countQuestionsByAllYes(groupAnswer string) int {
	questionsSeen := make([]rune, 0)
	for _, r := range groupAnswer {
		if r == ',' {
			continue
		}
		if !function.Contains(questionsSeen, r) {
			questionsSeen = append(questionsSeen, r)
		}
	}

	return len(questionsSeen)
}

func countQuestionsWhereAllAnsweredYes(groupAnswer string) int {

	numberOfPeople := len(strings.Split(groupAnswer, ","))
	distinctAnswerMap := make(map[rune]int, 0)

	for _, r := range groupAnswer {
		if r == ',' {
			continue
		}
		if val, ok := distinctAnswerMap[r]; ok {
			distinctAnswerMap[r] = val + 1
		} else {
			distinctAnswerMap[r] = 1
		}
	}

	total := 0

	if numberOfPeople == 1 {
		total += len(distinctAnswerMap)
	} else {
		for _, v := range distinctAnswerMap {
			if v == numberOfPeople {
				total++
			}
		}
	}

	return total
}
