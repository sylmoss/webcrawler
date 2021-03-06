package main

import (
    "fmt"
    "net/http" 
    "net/url"
    "io/ioutil"
    "sync"
)

type Crawler struct {
    crawled map[string]bool
    mux     sync.Mutex
}

func (c *Crawler) isCrawled(url string) bool {
    c.mux.Lock()
    defer c.mux.Unlock()
    
    _, ok := c.crawled[url]
    if ok {
        return true
    }
    c.crawled[url] = true

    return false
}

func (c *Crawler) Crawl(tabSpace string, url string) {
    var wg sync.WaitGroup

    if c.isCrawled(url) {
        return
    }

    fmt.Println(tabSpace + url)

    for _, childUrl  := range fetchUrlsFrom(url) {
        wg.Add(1)
        go func(childUrl string) {
            defer wg.Done()
            c.Crawl("---" + tabSpace, childUrl)
        }(childUrl)
    }

    wg.Wait()
    return
}

func fetchUrlsFrom(url string) []string {
    resp, err := http.Get(url) 
    if err == nil {
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)

        if err == nil {
            strBody := string(body)
            return extractUrlsFromHtml(url, strBody)
        }
    }
    return nil
}

func extractUrlsFromHtml(Url, body string) []string {
    regex := "(?s)<a[ t]+.*?href=\"((/).*?)\".*?>.*?</a>" 
    linkTags := NewFilter(regex).Find(body)
    var links []string
    baseUrl, _ := url.Parse(Url)
    if linkTags != nil {
        for _, z := range linkTags {
            url, err := url.Parse(z[1])
            if err == nil {
                links = append(links, baseUrl.ResolveReference(url).String()) 
            }
        }
    }
    return links
}
