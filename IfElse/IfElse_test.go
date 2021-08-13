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
		t.Fatalf("AssignIf returns incorrect result: expected %v, got %v", 14, ifelse.AssignInIf(5))
	}

	if ifelse.AssignInIf(0) != 0 {
		t.Fatalf("AssignIf returns incorrect result: expected %v, got %v", 0, ifelse.AssignInIf(0))
	}
}

func TestAssignInIfElse(t *testing.T) {
	if ifelse.AssignInIfElse(5) != 14 {
		t.Fatalf("AssignInIfElse returns incorrect result: expected %v, got %v", 14, ifelse.AssignInIfElse(5))
	}

	if ifelse.AssignInIfElse(0) != 0 {
		t.Fatalf("AssignInIfElse returns incorrect result: expected %v, got %v", 0, ifelse.AssignInIfElse(0))
	}
}

func TestSwitch(t *testing.T) {
	if ifelse.Switch(2) != "Even" {
		t.Fatalf("Switch(%v) returns incorrect result: expected %v, got %v", 2, "Even", ifelse.AssignInIfElse(2))
	}
	if ifelse.Switch(3) != "Odd" {
		t.Fatalf("Switch(%v) returns incorrect result: expected %v, got %v", 3, "Odd", ifelse.Switch(3))
	}
}
