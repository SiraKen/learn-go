package loop_test

import (
	"testing"

	"github.com/siraken/learn-golang/loop"
)

func TestLoop(t *testing.T) {
	expected := "1, 2, 3, 4"
	got := loop.CountNumbers(1, 4)

	if got != expected {
		t.Errorf("Loop() = %s; expected %s", got, expected)
	}
}
