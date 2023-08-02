package score_test

import (
	"testing"

	"github.com/siraken/learn-golang/score"
)

func TestIsPassed(t *testing.T) {
	expected := true
	got := score.IsPassed(85)

	if expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}

	expected = false
	got = score.IsPassed(79)

	if expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
}
