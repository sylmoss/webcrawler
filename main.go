package main

import "os"

func main() {
    host := os.Args[1]

    crawler := &Crawler{
        crawled: make(map[string]bool),
    }

    crawler.Crawl("->", host)
}
