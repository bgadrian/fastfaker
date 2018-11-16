package fastfaker

import (
	"testing"
)

func TestClampInt(t *testing.T) {
	if clampInt(1, 10, 20) != 10 {
		t.Errorf("failed for 1, 10, 20")
	}

	if clampInt(100, 10, 20) != 20 {
		t.Errorf("failed for 1, 10, 20")
	}
}
