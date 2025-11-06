package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebChecker(url string) bool {
	return url != "wss://net.com"
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"wss://net.com",
		"https://google.com",
		"https://github.com",
	}

	want := map[string]bool{
		"wss://net.com":      false,
		"https://google.com": true,
		"https://github.com": true,
	}

	got := CheckWebsite(mockWebChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func slowWebsiteStub(string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 1000)
	for i := range len(urls) {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsite(slowWebsiteStub, urls)
	}
}

func BenchmarkSlowCheckWebsite(b *testing.B) {
	urls := make([]string, 10)
	for i := range len(urls) {
		urls[i] = "a url"
	}

	for b.Loop() {
		SlowCheckWebsite(slowWebsiteStub, urls)
	}
}
