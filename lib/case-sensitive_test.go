package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	t.Run("入力例1", func(t *testing.T) {
		assert.Equal(t, "case-insensitive", CaseSensitive("AbC", "ABc"))
	})
	t.Run("入力例2", func(t *testing.T) {
		assert.Equal(t, "same", CaseSensitive("xyz", "xyz"))
	})
	t.Run("入力例3", func(t *testing.T) {
		assert.Equal(t, "different", CaseSensitive("aDs", "kjH"))
	})
}
