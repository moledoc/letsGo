package main

import (
	"fmt"
	"sync"
)

// self added:
type mutexURLCache struct {
	mux      sync.Mutex
	seenURLs map[string]bool
}

// self added: init the url cache struct
var urlCache = mutexURLCache{seenURLs: make(map[string]bool)}

// self added:
// method to check, if given url is cached.
// if yes, then return true (means corresponding goroutine will exit crawl from given url)
// if no, then add the url to cache and return false (goroutine will crawl this url)
func (urlCache *mutexURLCache) isSeen(url string) bool {
	urlCache.mux.Lock()
	defer urlCache.mux.Unlock()
	if urlCache.seenURLs[url] {
		return true
	}
	urlCache.seenURLs[url] = true
	return false
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func crawlInner(url string, depth int, fetcher Fetcher, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done() // self added
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	// self added: check if we have seen this url
	if urlCache.isSeen(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		waitGroup.Add(1) // self added
		go crawlInner(u, depth-1, fetcher, waitGroup)
		//fmt.Println(u)
		//fmt.Println((*urlCache).seenURLs)
	}
	return
}

// self added: create a wrapper function (for waiting group).
// without waiting group, the code will exit after the first goroutine finishes.
func Crawl(url string, depth int, fetcher Fetcher) {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)
	go crawlInner(url, depth, fetcher, waitGroup)
	waitGroup.Wait()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
