package main

import "testing"

func TestGetRowColumn(t *testing.T) {
	testCases := []struct {
		seat              string
		expectedRowNumber int
		expectedColNumber int
	}{
		{
			seat:              "FBBFFFFRLR",
			expectedRowNumber: 48,
			expectedColNumber: 5,
		},
		{
			seat:              "FFBFFBFRRR",
			expectedRowNumber: 18,
			expectedColNumber: 7,
		},
		{
			seat:              "BFFBBFFLRL",
			expectedRowNumber: 76,
			expectedColNumber: 2,
		},
		{
			seat:              "BBFBBBBLLL",
			expectedRowNumber: 111,
			expectedColNumber: 0,
		},
		{
			seat:              "BFBBFFFLLR",
			expectedRowNumber: 88,
			expectedColNumber: 1,
		},
		{
			seat:              "BBFFBFFRRL",
			expectedRowNumber: 100,
			expectedColNumber: 6,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.seat, func(t *testing.T) {
			row := getRowNumber(testCase.seat)
			col := getColumnNumber(testCase.seat)

			if row != testCase.expectedRowNumber {
				t.Fatalf("Incorrect row number: %s, %d. Result was: %d", testCase.seat, testCase.expectedRowNumber, row)
			}

			if col != testCase.expectedColNumber {
				t.Fatalf("Incorrect col number: %s, %d. Result was: %d", testCase.seat, testCase.expectedColNumber, col)
			}
		})
	}
}
