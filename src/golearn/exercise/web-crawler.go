package exercise

import (
	"fmt"
	"sync"
)

// Fetcher 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// Cache cache for fetch
type Cache struct {
	fetched map[string]bool
	mux     sync.Mutex
}

var cache = Cache{fetched: make(map[string]bool)}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	if depth <= 0 {
		return
	}

	cache.mux.Lock()
	_, ok := cache.fetched[url]
	if ok == true {
		cache.mux.Unlock()
		return
	}
	cache.fetched[url] = true
	cache.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	return
}

// FakeFetcher 是返回若干结果的 Fetcher。
type FakeFetcher map[string]*FakeResult

// FakeResult 结果
type FakeResult struct {
	Body string
	Urls []string
}

// Fetch 实现方法
func (f FakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.Body, res.Urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}
