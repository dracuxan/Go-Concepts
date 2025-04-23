package main

import "testing"

func TestHello(t *testing.T) {
	orignal := helloWorld("Nisarg")
	expected := "Hello, Nisarg!"

	if orignal != expected {
		t.Errorf("orignal: '%v' expected: '%v'", orignal, expected)
	}
}

func BenchmarkHello(b *testing.B) {
	for b.Loop() {
		helloWorld("Nisarg")
	}
}
