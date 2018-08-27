package main

import (
    "fmt"
    "net/http" 
    "net/url"
    "regexp"
    "io/ioutil"
    "os"
    "sync"
)

type Crawler struct {
    crawled map[string]bool
    mux     sync.Mutex
}

func (c *Crawler) visit(url string) bool {
    c.mux.Lock()
    defer c.mux.Unlock()
    
    _, ok := c.crawled[url]
    if ok {
        return true
    }
    c.crawled[url] = true

    return false
}

func (c *Crawler) Crawl(url string) {
    var wg sync.WaitGroup

    v := c.visit(url)
    if v {
        return
    }

    for _, u  := range fetch(url) {
        wg.Add(1)

        go func(u string) {
            defer wg.Done()
            c.Crawl(u)
        }(u)
    }

    wg.Wait()
    return
}

func main() {
    host := os.Args[1]
    crawler := &Crawler{
        crawled: make(map[string]bool),
    }
    crawler.Crawl(host)
}

func fetch(url string) []string {
    resp, err := http.Get(url) 
    if err != nil {
        fmt.Println("not found: %s", url)
    } else {
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)

        if err != nil {
            fmt.Println("Read error has occured")
        } else {
            strBody := string(body)
            return extractUrls(url, strBody)
        }
    }
    return nil
}

func extractUrls(Url, body string) []string {
    newUrls := regexp.MustCompile("(?s)<a[ t]+.*?href=\"((/).*?)\".*?>.*?</a>").FindAllStringSubmatch(body, -1)
    var links []string
    baseUrl, _ := url.Parse(Url)
    if newUrls != nil {
        for _, z := range newUrls {
            ur, err := url.Parse(z[1])
            if err == nil {
                links = append(links, baseUrl.ResolveReference(ur).String()) 
            }
        }
    }
    return links
}