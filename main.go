package main

import (
	"errors"
	"fmt"
	"time"
)

// ErrYearNotGrigorian ...
var ErrYearNotGrigorian = errors.New("not a Grigorian year")

func IsLeapYear(year int) bool {

	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	}

	return year%4 == 0
}

func TryIsLeapYear(year int) (bool, error) {
	if year < 1582 {
		return false, ErrYearNotGrigorian
	}

	return IsLeapYear(year), nil
}

func main() {
	year := time.Now().Year()
	printYear(year)

	year = inputGrigorianYear(inputInteger, isGrigorianYear)
	printYear(year)
}

func printYear(year int) {
	fmt.Print("Year ", year, " is ")
	if !IsLeapYear(year) {
		fmt.Print("not ")
	}
	fmt.Println("leap")
}

func inputInteger() int {
	var value int
	for {
		if _, err := fmt.Scanln(&value); err == nil {
			return value
		}
	}
}

func isGrigorianYear(y int) bool {
	return y >= 1582
}

func inputGrigorianYear(input func() int, validator func(int) bool) int {
	
	for {
		year := input()
		if validator(year) {
			return year
		}
	}
}
