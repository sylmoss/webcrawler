package main

import (
    "testing"
)

func TestExtractUrlsFromHtmlBody(t *testing.T) {
    body := buildBody()
    url := "https://monzo.com/"
    
    urls := extractUrlsFromHtml(url, body)

    if len(urls) != 6 {
        t.Error("Tags were not right filtered from body html")
    }

    if isAllUrlsExtracted(urls) != true {
        t.Error("Tags were not right filtered from body html")
    } 
}

func isAllUrlsExtracted(filteredTags []string) bool {
    return (filteredTags[0] == "https://monzo.com/") &&
    (filteredTags[1] == "https://monzo.com/about") &&
    (filteredTags[2] == "https://monzo.com/blog") &&
    (filteredTags[3] == "https://monzo.com/community") &&
    (filteredTags[4] == "https://monzo.com/faq") &&
    (filteredTags[5] == "https://monzo.com/download")
}