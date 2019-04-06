练习：Web 爬虫  

在这个练习中，我们将会使用 Go 的并发特性来并行化一个 Web 爬虫。

修改 Crawl 函数来并行地抓取 URL，并且保证不重复。

提示：你可以用一个 map 来缓存已经获取的 URL，但是要注意 map 本身并不是并发安全的！  

```go
package main

import (
    "fmt"
    "sync"
)

type Fetcher interface {
    Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {

    defer wg.Done()

    if depth <= 0 {
        return
    }

    if cache.has(url) {
        // fmt.Println("already exist.")
        return
    }

    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("found: %s %q\n", url, body)

    for _, u := range urls {
        wg.Add(1)
        go Crawl(u, depth-1, fetcher, wg)
    }

    return
}

type Cache struct {
    cache map[string]bool
    mutex sync.Mutex
}

func (cache *Cache) add(url string) {
    cache.mutex.Lock()
    cache.cache[url] = true
    cache.mutex.Unlock()
}

func (cache *Cache) has(url string) bool {
    cache.mutex.Lock()
    defer cache.mutex.Unlock()
    _, ok := cache.cache[url]
    if !ok {
        cache.cache[url] = true
    }
    return ok
}

var cache Cache = Cache{
    cache: make(map[string]bool),
}

func main() {
    var wg sync.WaitGroup

    wg.Add(1)
    go Crawl("https://golang.org/", 4, fetcher, &wg)

    wg.Wait()
}

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
            "https://golang.org/pkg/os1/",
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
```

输出：  
```go
found: https://golang.org/ "The Go Programming Language"
not found: https://golang.org/cmd/
found: https://golang.org/pkg/ "Packages"
not found: https://golang.org/pkg/os1/
found: https://golang.org/pkg/fmt/ "Package fmt"
found: https://golang.org/pkg/os/ "Package os"

Program exited.
```
