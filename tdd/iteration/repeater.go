package iteration

const repeatCount = 5

func Repeat(c string) (repeated string) {
	for range repeatCount {
		repeated += c
	}
	return
}
