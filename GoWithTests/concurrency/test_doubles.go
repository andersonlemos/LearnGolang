package concurrency

import "time"

func mockVerifyWebsite(url string) bool {
	if url == "https://amazon.com" {
		return false
	}

	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
