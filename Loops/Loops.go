package loops

// golang's for loop
func ForLoop(endpoint int) int {
	var sum int
	for i := 0; i < endpoint; i++ {
		sum += i
	}
	return sum
}

// golang's while loop
func BasicallyWhileLoop(endpoint int) int {
	sum := 0
	for sum < endpoint {
		sum += 1
	}
	return sum
}
