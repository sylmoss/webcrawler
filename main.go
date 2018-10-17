package main

import (
    "os"
    "fmt"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please, provide the host url as an argument.")
    } else {
        host := os.Args[1] 

        crawler := &Crawler{
            crawled: make(map[string]bool),
        }

        crawler.Crawl("->", host)
    }
}
