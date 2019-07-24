package binSearch

import (
//	"errors"
	"testing"
)

var branchInsertTests = []struct {
	in1 int
	in2 interface{}
	out error
}{
	{1, "test1", nil},
	{3, "test3", nil},
	{15, "test15", nil},
	{10, "test10", nil},
	{8, "test8", nil},
	{9, "test9", nil},
	{10, "test10new", nil},
}

var branchFindTests = []struct {
	in   int
	out1 interface{}
	out2 string
}{
	{9, "test9", ""},
	{29, nil, "Address not found"},
	{10, "test10new", ""},
	{8, "test8", ""},
	{1, "test1", ""},
	{33, nil, "Address not found"},
	{3, "test3", ""},
	{15, "test15", ""},
}

func TestNew(t *testing.T) {
	got := New("Nur")
	if got == nil {
		t.Error("New('Nur') ")
	}
}

func TestInsert(t *testing.T) {
	got := New("Nur")
	for _, tt := range branchInsertTests {
		b := got.Insert(tt.in1, tt.in2)
		if b != tt.out {
			t.Errorf("Half(%d, %d) => %d, want %d", tt.in1, tt.in1, b, tt.out)
		}
	}
}

func TestFind(t *testing.T) {
	got := New("Nur")
	for _, tt := range branchInsertTests {
		_ = got.Insert(tt.in1, tt.in2)
	}

	for _, tt := range branchFindTests {
		b, e := got.Find(tt.in)
		if b != tt.out1 {
			t.Errorf("Half(%d) => %d, want %d", tt.in, b, tt.out1)
			if e != nil {
				t.Errorf("Half(%d) => %s, want nil", tt.in, e)
			}
		}
		if e != nil {
			if e.Error() != tt.out2 {
				t.Errorf("Half(%d) => %s, want %s", tt.in, e, tt.out2)
			}
		}
	}
}
