package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadNumber(t *testing.T) {
	n := 5
	lines := []string{
		"..#..",
		".##..",
		"..#..",
		"..#..",
		".###.",
	}

	assert.Equal(t, 1, ReadNumber(lines, n))
}
