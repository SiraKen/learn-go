package greet_test

import (
	"testing"

	"github.com/siraken/learn-golang/greet"
)

func TestEnglish(t *testing.T) {
	expected := "Hello, Go"
	got := greet.English("Go")

	if expected != got {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestJapanese(t *testing.T) {
	expected := "こんにちは、Go"
	got := greet.Japanese("Go")

	if expected != got {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
