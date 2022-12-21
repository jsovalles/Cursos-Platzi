package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("AddSuccess", func(t *testing.T) {
		c := require.New(t)
		result := Add(20, 2)
		expected := 22
		c.Equal(result, expected)
	})
}
