package dependencies

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nisarg")

	got := buffer.String()
	want := "Hello, Nisarg"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
