package ifelse_test

import (
	"testing"

	ifelse "github.com/moledoc/letsgo/IfElse"
)

func TestBasicIf(t *testing.T) {
	if ifelse.BasicIf(5, 3) != true {
		t.Fatalf("BasicIf returns incorrect result: expected %v, got %v", true, ifelse.BasicIf(5, 3))
	}
}

func TestAssignIf(t *testing.T) {
	if ifelse.AssignInIf(5) != 14 {
		t.Fatalf("TestAssignIf returns incorrect result: expected %v, got %v", 14, ifelse.AssignInIf(5))
	}

	if ifelse.AssignInIf(0) != 0 {
		t.Fatalf("TestAssignIf returns incorrect result: expected %v, got %v", 0, ifelse.AssignInIf(0))
	}
}

func TestAssignInIfElse(t *testing.T) {
	if ifelse.AssignInIfElse(5) != 14 {
		t.Fatalf("TestAssignInIfElse returns incorrect result: expected %v, got %v", 14, ifelse.AssignInIfElse(5))
	}

	if ifelse.AssignInIfElse(0) != 0 {
		t.Fatalf("TestAssignInIfElse returns incorrect result: expected %v, got %v", 0, ifelse.AssignInIfElse(0))
	}
}
