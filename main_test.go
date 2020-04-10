package main

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("YamadaTaro")
	expected := "Hello, YamadaTaro"
	if result != expected {
		t.Errorf("Result:%v\nTrue:%v", result, expected)
	}
}
