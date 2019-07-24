package binSearch

import (
	"errors"
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
	out2 error
}{
	{9, "test9", nil},
	{29, nil, errors.New("Address not found")},
	{10, "test10new", nil},
	{8, "test8", nil},
	{1, "test1", nil},
	{33, "nil", errors.New("Address not found")},
	{3, "test3", nil},
	{15, "test15", nil},
}

func TestNew(t *testing.T) {
	got := New("Nur")
	if got == nil {
		t.Error("New('Nur') ")
	}
}

func (got *Cache) TestInsert(t *testing.T) {
	for _, tt := range branchInsertTests {
		b := got.Insert(tt.in1, tt.in2)
		if b != tt.out {
			t.Errorf("Half(%d, %d) => %d, want %d", tt.in1, tt.in1, b, tt.out)
		}
	}
}

func (got *Cache) TestFind(t *testing.T) {
	for _, tt := range branchFindTests {
		b, e := got.Find(tt.in)
		if b != tt.out1 {
			t.Errorf("Half(%d) => %d, want %d", tt.in, b, tt.out1)
		}
		if e != tt.out2 {
			t.Errorf("Half(%d) => %d, want %d", tt.in, b, tt.out2)
		}
	}
}
