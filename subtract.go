package coverage

func SubtractPositives(a, b int) int {
	if a < 0 || b < 0 {
		return -1
	}

	return a - b
}

func SubtractNegatives(a, b int) int {
	if a < 0 || b < 0 {
		return 1
	}

	return a - b
}
