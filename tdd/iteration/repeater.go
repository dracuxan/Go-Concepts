package iteration

func Repeat(c string, count int) (repeated string) {
	for range count {
		repeated += c
	}
	return
}
