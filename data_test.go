package main

import "testing"
import "fmt"

func TestCells(t *testing.T) {
	if len(*Cells) != 81 {
		t.Error(fmt.Sprintf("TestCells - expected 81 but got %d", len(*Cells)))
	}
}

func TestCross(t *testing.T) {
	var unit *UnitType
	unit = Cross("a", "b")
	if (*unit)["ab"] != true {
		t.Error("TestCross - expected ab")
	}
	unit = Cross("abc", "123")
	if (*unit)["a1"] != true || (*unit)["a2"] != true || (*unit)["c3"] != true {
		t.Error("TestCross - expected ab")
	}
}
