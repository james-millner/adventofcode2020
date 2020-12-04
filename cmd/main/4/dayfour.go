package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/james-millner/adventofcode2020/pkg/file"
)

var validFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func main() {
	box := packr.NewBox("./input")
	batchPassportFile := file.RetreiveInputFileAsString("dayfour.txt", box)

	regex := regexp.MustCompile(`(?m)^\s*$`).ReplaceAllString(strings.TrimSpace(batchPassportFile), "-")
	passports := strings.Split(regex, "-")

	formattedPassports := Map(passports, createPassportString)

	validPassports := filter(formattedPassports, passportVerifier)

	fmt.Printf("%d\n", len(validPassports))
}

func passportVerifier(passport string) bool {

	if passport == "" {
		return false
	}

	numberOfValidFields := 0

	for _, validField := range validFields {
		if strings.Contains(passport, validField) {
			numberOfValidFields++
		}
	}

	if numberOfValidFields != len(validFields) {
		return false
	}

	return true
}

func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func createPassportString(passport string) string {
	pieces := strings.Split(passport, "\n")
	return strings.Join(pieces, " ")
}
