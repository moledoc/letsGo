package ifelse

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
