package moretypes_test

import (
	"testing"

	moretypes "github.com/moledoc/letsgo/MoreTypes"
)

func TestSubtractr(t *testing.T) {
	subtractor := moretypes.Subtractr()
	val := subtractor(5)
	if val != -5 {
		t.Fatalf("moretypes.Subtractr works incorrectly: expected %v, got %v", -5, val)
	}
	val = subtractor(2)
	if val != -7 {
		t.Fatalf("moretypes.Subtractr works incorrectly: expected %v, got %v", -7, val)
	}
	val = subtractor(2)
	if val != -9 {
		t.Fatalf("moretypes.Subtractr works incorrectly: expected %v, got %v", -9, val)
	}
}
