package concurrency

type WebSiteChecker func(string) bool
type Result struct {
	string
	bool
}

func VerifyWebsites(wc WebSiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	channelResult := make(chan Result, len(urls))

	for _, url := range urls {
		go func(u string) {
			channelResult <- Result{u, wc(u)}

		}(url)
		result := <-channelResult
		results[result.string] = result.bool
	}
	return results
}
