package main

import "testing"

func TestIsPasswordCandidate(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{123447, true},
		{123455, true},
		{223450, false},
		{123789, false},

		{112233, true},
		{123444, false},
		{122444, true},
		{111122, true},
		{111122, true},
		{111122, true},
		{223334, true},
	}
	for _, tt := range tests {
		if got := IsPasswordCandidate(tt.input); got != tt.want {
			t.Errorf("IsPasswordCandidate(%d) = %v, want %v", tt.input, got, tt.want)
		}

	}
}
