package concurrency

type (
	WebsiteChecker func(string) bool
	result         struct {
		string
		bool
	}
)

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	resChan := make(chan result)
	for _, url := range urls {
		go func() {
			resChan <- result{url, wc(url)}
		}()
	}

	for range len(urls) {
		r := <-resChan
		results[r.string] = r.bool
	}

	return results
}

func SlowCheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
