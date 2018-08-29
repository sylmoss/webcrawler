package main

import (
    "testing"
)

func TestFilteBodyToFindLinkTags(t *testing.T) {
    body := buildBody()
    regex := "(?s)<a[ t]+.*?href=\"((/).*?)\".*?>.*?</a>" 
    filter := NewFilter(regex)

    filteredTags := filter.Find(body)

    if len(filteredTags) != 6 {
        t.Error("Tags were not right filtered from body html")
    }

    if isAllTagsFiltered(filteredTags) != true {
        t.Error("Tags were not right filtered from body html")
    } 
}

func isAllTagsFiltered(filteredTags [][]string) bool {
    return (filteredTags[0][1] == "/") &&
    (filteredTags[1][1] == "/about") &&
    (filteredTags[2][1] == "/blog") &&
    (filteredTags[3][1] == "/community") &&
    (filteredTags[4][1] == "/faq") &&
    (filteredTags[5][1] == "/download")
}