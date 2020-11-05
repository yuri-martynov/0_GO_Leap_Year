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

func TestIsNotLeap1580Year(t *testing.T) {
	b := IsLeapYear(1580)
	if b {
		t.Error("Year 1580 is before a Grigorian calendar")
	}
}
