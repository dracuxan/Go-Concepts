package dependencies

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(b io.Writer, name string) {
	fmt.Fprintf(b, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World!")
}
