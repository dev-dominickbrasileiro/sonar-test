package coverage

import "testing"

func TestSumPositives(t *testing.T) {
	if SumPositives(5, 3) != 8 {
		t.Error("result is different from expected")
	}

	if SumPositives(5, -3) != -1 {
		t.Error("negative check failed")
	}

	if SumPositives(-5, 3) != -1 {
		t.Error("negative check failed")
	}

	if SumPositives(-5, -3) != -1 {
		t.Error("negative check failed")
	}
}
