package coverage

import "testing"

func TestSubtractPositives(t *testing.T) {
	if SubtractPositives(5, 3) != 2 {
		t.Error("result is different from expected")
	}

	if SubtractPositives(5, -3) != -1 {
		t.Error("negative check failed")
	}

	if SubtractPositives(-5, 3) != -1 {
		t.Error("negative check failed")
	}

	if SubtractPositives(-5, -3) != -1 {
		t.Error("negative check failed")
	}
}

func TestSubtractNegatives(t *testing.T) {
	if SubtractNegatives(-5, -3) != -2 {
		t.Error("result is different from expected")
	}

	if SubtractNegatives(-5, 3) != 1 {
		t.Error("positive check failed")
	}

	if SubtractNegatives(5, -3) != 1 {
		t.Error("positive check failed")
	}

	if SubtractNegatives(5, 3) != 1 {
		t.Error("positive check failed")
	}
}
