package main

import "testing"

func TestHello(t *testing.T) {
	orignal := helloWorld("Nisarg")
	expected := "Hello Nisarg!"

	if orignal != expected {
		t.Errorf("got %v expected %v", orignal, expected)
	}
}
