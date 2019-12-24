package concurrency

type SiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckSites(checker SiteChecker, urls []string) map[string]bool  {
	results := map[string]bool{}
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, checker(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results;
}
