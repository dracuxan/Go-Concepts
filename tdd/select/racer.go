package selec

import (
	"net/http"
	"time"
)

func Racer(a string, b string) string {
	starta := time.Now()
	http.Get(a)
	aDuration := time.Since(starta)

	startb := time.Now()
	http.Get(b)
	bDuration := time.Since(startb)

	if aDuration < bDuration {
		return a
	}
	return b
}
