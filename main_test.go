package main

import "testing"

func TestIsLeap2000Year(t *testing.T) {
	b := IsLeapYear(2000)
	if !b {
		t.Errorf("expect true, got %t", b)
	}
}

func TestIsNotLeap2100Year(t *testing.T) {
	b := IsLeapYear(2100)
	if b {
		t.Errorf("expected false, got %t", b)
	}
}

func TestIsLeap2020Year(t *testing.T) {
	b := IsLeapYear(2020)
	if !b {
		t.Error("Year 2020 should be leap")
	}
}

func TestIsLeap2016Year(t *testing.T) {
	b := IsLeapYear(2016)
	if !b {
		t.Error("Year 2016 should be leap")
	}
}

func Test1580IsNotGrigorian(t *testing.T) {
	_, err := TryIsLeapYear(1580)
	if err != ErrYearNotGrigorian {
		t.Error("1580 is not a Grigorian year")
	}
}

func Test1581IsGrigorian(t *testing.T) {
	_, err := TryIsLeapYear(1581)
	if err != ErrYearNotGrigorian {
		t.Error("1581 is not a Grigorian year")
	}
}

func Test2020IsGrigorian(t *testing.T) {
	_, err := TryIsLeapYear(2020)
	if err == ErrYearNotGrigorian {
		t.Error("2020 is a Grigorian year")
	}
}

func Test2020IsLeap(t *testing.T) {
	b, _ := TryIsLeapYear(2020)
	if !b {
		t.Error("2020 should be Leap year")
	}
}

func Test2016IsLeap(t *testing.T) {
	b, _ := TryIsLeapYear(2016)
	if !b {
		t.Error("2016 should be Leap year")
	}
}

func Test2015IsGrigorianYear(t *testing.T) {
	b := isGrigorianYear(2015)
	if !b {
		t.Error("2015 is not GrigorianYear")
	}
}

func Test2017IsGrigorianYear(t *testing.T) {
	b := isGrigorianYear(2017)
	if !b {
		t.Error("2017 is not GrigorianYear")
	}
}

func Test1500IsNotGrigorianYear(t *testing.T) {
	b := isGrigorianYear(1500)
	if b {
		t.Error("1500 is not GrigorianYear")
	}
}

func Test2021ValidtorInput(t *testing.T) {

	var y int
	input := func() int { return 2021 }
	validator := func(year int) bool {
		y = year
		return true
	}

	_ = inputGrigorianYear(input, validator)
	if y != 2021 {
		t.Error("Validator shoud have an input 2021")
	}
}
