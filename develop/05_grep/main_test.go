package main

import "testing"

func TestGrep(t *testing.T) {
	type inputFlags struct {
		after      int
		before     int
		context    int
		count      bool
		ignoreCase bool
		inversion  bool
		lineNum    bool
		pattern    string
	}
}
