package chapter3_test

import (
	"testing"

	"ddd-golang/chapter3"
)

func Test_Point(t *testing.T) {
	a := chapter3.NewPoint(1, 1)
	b := chapter3.NewPoint(1, 1)
	if a != b {
		t.Fatal("a and  b were not equal")
	}
}
