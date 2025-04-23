package iteration

func Repeat(c string) (repeated string) {
	for i := 5; i > 0; i-- {
		repeated += c
	}
	return
}
