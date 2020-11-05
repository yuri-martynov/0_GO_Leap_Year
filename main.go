package main

func main() {}

func IsLeapYear(year int) bool {

	switch {
	case year < 1582:
		return false
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	}
	/*
		if year < 1582 {
			return false
		} else if year%400 == 0 {
			return true
		} else if year%100 == 0 {
			return false
		}
	*/
	return year%4 == 0
}
