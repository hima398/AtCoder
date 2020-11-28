package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullSearch(t *testing.T) {
	tests := []struct {
		name          string
		a, b, c, x, y int
		expected      int
	}{
		{"Example 1",
			1500, 2000, 1600, 3, 2,
			7900},
		{"Example 2",
			1500, 2000, 1900, 3, 2,
			8500},
		{"Example 3",
			1500, 2000, 500, 90000, 100000,
			100000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, FullSearch(tt.a, tt.b, tt.c, tt.x, tt.y))
		})
	}
}

func TestO1(t *testing.T) {
	tests := []struct {
		name          string
		a, b, c, x, y int
		expected      int
	}{
		{"Example 1",
			1500, 2000, 1600, 3, 2,
			7900},
		{"Example 2",
			1500, 2000, 1900, 3, 2,
			8500},
		{"Example 3",
			1500, 2000, 500, 90000, 100000,
			100000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, O1(tt.a, tt.b, tt.c, tt.x, tt.y))
		})
	}
}
