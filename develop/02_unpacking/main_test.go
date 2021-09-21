package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpacking(t *testing.T) {
	table := []struct{
		in string
		out string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
	}

	for _, val := range table {
		result := Unpacking(val.in)

		assert.Equal(t, result, val.out)
	}
}
