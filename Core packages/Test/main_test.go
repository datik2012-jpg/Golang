package main

import "testing"

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld()
	expected := "hello world"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
