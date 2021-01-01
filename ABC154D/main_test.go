package main

import (
	"testing"
)

func TestExpectedValue(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"1", args{1}, 1.0},
		{"2", args{2}, 1.5},
		{"3", args{3}, 2.0},
		{"4", args{4}, 2.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpectedValue(tt.args.n); got != tt.want {
				t.Errorf("ExpectedValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
