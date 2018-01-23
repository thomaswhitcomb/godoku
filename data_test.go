package main

import "testing"
import "fmt"

func TestCells(t *testing.T) {
	if len(*Cells) != 81 {
		t.Error(fmt.Sprintf("TestCells - expected 81 but got %d", len(*Cells)))
	}
}

func TestCross(t *testing.T) {
	tests := []struct {
		name    string
		parm1   string
		parm2   string
		size    int
		indexes []string
	}{
		{"1", "a", "b", 1, []string{"ab"}},
		{"2", "abc", "123", 9, []string{"a1", "a2", "c3", "b1", "b2", "b3", "c1", "c2", "c3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unit := Cross(tt.parm1, tt.parm2)
			for _, x := range tt.indexes {
				if len(*unit) != tt.size {
					t.Errorf("bad size. expected %d and got %d", tt.size, len(*unit))
				}
				if (*unit)[x] != true {
					t.Errorf("index %s should be true and it is not", x)
				}
			}
		})
	}
}
