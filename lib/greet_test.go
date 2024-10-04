package lib

import (
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		what string
	}{
		{"", "Hello, anonymous!"},
		{"Max", "Hello, Max!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.name); got != tt.what {
				t.Errorf("Greet() = %v, want %v", got, tt.what)
			}
		})
	}
}
