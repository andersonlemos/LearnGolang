package concurrency

import (
	"reflect"
	"testing"
)

func BenchmarkVerifyWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "one url"
	}
	for n := 0; n < b.N; n++ {
		VerifyWebsites(slowStubWebsiteChecker, urls)
	}
}

func TestVerifyWebsite(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
	}
	expected := map[string]bool{
		"https://google.com":   true,
		"https://facebook.com": true,
		"https://golang.org":   true,
		"https://amazon.com":   false,
	}

	result := VerifyWebsites(mockVerifyWebsite, websites)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Verify Websites failed: expected %v, got %v", expected, result)
	}

}
