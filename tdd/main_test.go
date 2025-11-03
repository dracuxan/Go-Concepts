package main

import (
	"testing"
)

func TestMain(m *testing.T) {
	want := "Nisarg"
	got := getName()

	if want != got {
		m.Errorf("got %s want %s", got, want)
	}
}
