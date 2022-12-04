package coverage

func SumPositives(a, b int) int {
	if a < 0 || b < 0 {
		return -1
	}

	return a + b
}

func SumNegatives(a, b int) int {
	if a < 0 || b < 0 {
		return 1
	}

	return a + b
}
