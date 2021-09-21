package main

import "testing"

func TestHello(t *testing.T) {
	//var str, result = "a4bc2d5e", "aaaabccddddde"
	//realResult := Unpacking(str)
	//
	//if realResult != result {
	//	t.Errorf("expected result: %s != %s real result", result, realResult )
	//}

	var x, y, result = 2, 2, 4

	realResult := Hello(x, y)

	if realResult != result {
		t.Errorf("%d != %d")
	}
}
