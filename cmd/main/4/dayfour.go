package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/james-millner/adventofcode2020/pkg/file"
	"github.com/james-millner/adventofcode2020/pkg/function"
)

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

var optionalFields = []string{
	"cid",
}

var hairColours = []string{
	"amb",
	"blu",
	"brn",
	"gry",
	"grn",
	"hzl",
	"oth",
}

func main() {
	box := packr.NewBox("./input")
	batchPassportFile := file.RetreiveInputFileAsString("dayfour.txt", box)

	regex := regexp.MustCompile(`(?m)^\s*$`).ReplaceAllString(strings.TrimSpace(batchPassportFile), "-")
	passports := strings.Split(regex, "-")

	formattedPassports := function.Map(passports, createPassportString)

	// for _, v := range formattedPassports {
	// 	fmt.Printf("%s\n", extendedPassportVerifier(v))
	// }

	validPassportsWithBasicValidation := function.Filter(formattedPassports, basicPassportVerifier)
	fmt.Printf("Valid passports with basic validation: %d\n", len(validPassportsWithBasicValidation))

	validPassportsWithExtendedValidation := function.Filter(formattedPassports, extendedPassportVerifier)
	fmt.Printf("Valid passports with extended validation: %d\n", len(validPassportsWithExtendedValidation))
}

func basicPassportVerifier(passport string) bool {

	if passport == "" {
		return false
	}

	numberOfValidFields := 0

	for _, validField := range requiredFields {
		if strings.Contains(passport, validField) {
			numberOfValidFields++
		}
	}

	if numberOfValidFields != len(requiredFields) {
		return false
	}

	return true
}

func extendedPassportVerifier(passport string) bool {

	if !basicPassportVerifier(passport) {
		return false
	}

	passportSlice := strings.Split(passport, ",")

	passportMap := make(map[string]string)
	// put slice values into map
	for _, piece := range passportSlice {
		keyValue := strings.Split(piece, ":")
		passportMap[strings.TrimSpace(keyValue[0])] = strings.TrimSpace(keyValue[1])
	}

	if val, ok := passportMap["byr"]; ok {
		byr, _ := strconv.Atoi(val)
		if byr < 1920 || byr > 2002 {
			return false
		}
	}

	if val, ok := passportMap["iyr"]; ok {
		iyr, _ := strconv.Atoi(val)
		if iyr < 2010 || iyr > 2020 {
			return false
		}
	}

	if val, ok := passportMap["eyr"]; ok {
		expr, _ := regexp.Compile("([0-9]{4})")

		if !expr.MatchString(val) {
			return false
		}

		eyr, _ := strconv.Atoi(val)
		if eyr < 2020 || eyr > 2030 {
			return false
		}
	}

	if val, ok := passportMap["hgt"]; ok {
		heightAndMetric := regexp.MustCompile("([0-9]+)([a-z]+)").FindStringSubmatch(val)

		if len(heightAndMetric) == 0 {
			return false
		}

		heightValue := heightAndMetric[1]
		heightMetric := heightAndMetric[2]

		if heightMetric == "cm" {
			val, _ := strconv.Atoi(heightValue)
			if val < 150 || val > 193 {
				return false
			}
		} else if heightMetric == "in" {
			val, _ := strconv.Atoi(heightValue)
			if val < 59 || val > 76 {
				return false
			}
		}

		if heightMetric != "cm" && heightMetric != "in" {
			return false
		}
	}

	if val, ok := passportMap["hcl"]; ok {
		expr, _ := regexp.Compile("(#[0-9a-f]{6})")
		if !expr.MatchString(val) {
			return false
		}
	}

	if val, ok := passportMap["ecl"]; ok {
		if !function.Find(hairColours, val) {
			return false
		}
	}

	if val, ok := passportMap["pid"]; ok {
		_, err := regexp.MatchString(`(\d{9})`, val)
		if err != nil || len(val) != 9 {
			return false
		}

	}

	return true
}

func createPassportString(passport string) string {
	pieces := strings.FieldsFunc(passport, Split)

	return strings.Join(pieces, ",")
}

func Split(r rune) bool {
	return r == '\n' || r == ' '
}
