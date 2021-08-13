package ifelse

import "math"

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
