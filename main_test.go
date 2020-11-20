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