package toml_parser

import (
	"testing"
)

func TestValidateTOML(t *testing.T) {
	s := "corrext toml"

	got := ValidateTOML(s)
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
