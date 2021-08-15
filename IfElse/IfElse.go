package ifelse

import (
	"fmt"
	"math"
	"time"
)

// basic if statement
func BasicIf(a int, b int) bool {
	if a > b {
		return true
	}
	return false
}

// if statement with init statement/assign variable in it
func AssignInIf(a int) int {
	if b := 3*a - 1; a < b {
		return b
	}
	return a
}

// if-else statement
func AssignInIfElse(a int) int {
	if b := 3*a - 1; a > b {
		return a
	} else {
		return b
	}
}

// switch
func Switch(a float64) string {
	switch modA := math.Mod(a, 2); modA {
	case 0:
		return "Even"
	case 1:
		return "Odd"
	default:
		return ""
	}
}

// you can use functions inside a switch statement
func In6Days(today time.Weekday) time.Weekday {
	return today + 6
}

// switch v2
func SwitchV2() string {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		return "Today."
	case today + 1:
		return "Tomorrow."
	case In6Days(today):
		return fmt.Sprintf("In %v days.\n", 6)
	default:
		return "Between 2-5 days"
	}
}
