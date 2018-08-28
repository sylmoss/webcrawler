package main

import (
    "fmt"
    "testing"
)

func TestFileBodyToFindLinkTags(t *testing.T) {
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

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
  if a == b {
    return
    }
    if len(message) == 0 {
        message = fmt.Sprintf("%v != %v", a, b)
    }
    t.Fatal(message)
}

func isAllTagsFiltered(filteredTags [][]string) bool {
    return (filteredTags[0][1] == "/") &&
    (filteredTags[1][1] == "/about") &&
    (filteredTags[2][1] == "/blog") &&
    (filteredTags[3][1] == "/community") &&
    (filteredTags[4][1] == "/faq") &&
    (filteredTags[5][1] == "/download")
}

func buildBody() string {
    return "<header class=\"c-header\" id=\"navigation\">" +
              "<div class=\"c-header__container\">" +
                "<a href=\"/\" class=\"c-header__logo\" title=\"Monzo home page\">Monzo</a>" +
                "<nav class=\"c-header__nav\" aria-labelledby=\"menu-toggle\">" +
                  "<button class=\"c-header__toggle\" id=\"menu-toggle\" aria-expanded=\"false\" title=\"menu\" aria-controls=\"menu\">" +
                    "<span class=\"c-header__open\">Menu <span aria-hidden=\"true\">≡</span></span>" +
                    "<span class=\"c-header__close\">Close menu <span aria-hidden=\"true\"><svg width=\"10\" height=\"10\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M5 3.586L8.182.404a.5.5 0 0 1 .707 0l.707.707a.5.5 0 0 1 0 .707L6.414 5l3.182 3.182a.5.5 0 0 1 0 .707l-.707.707a.5.5 0 0 1-.707 0L5 6.414 1.818 9.596a.5.5 0 0 1-.707 0L.404 8.89a.5.5 0 0 1 0-.707L3.586 5 .404 1.818a.5.5 0 0 1 0-.707L1.11.404a.5.5 0 0 1 .707 0L5 3.586z\" fill=\"#FFF\" fill-rule=\"evenodd\"></path></svg></span></span>" +
                  "</button>" +
                  "<div class=\"c-header__menu\" id=\"menu\" aria-hidden=\"true\">" +
                    "<ul class=\"c-header__list\" role=\"menu\">" +
                      "<li class=\"c-header__item\" role=\"menuitem\"><a href=\"/about\" class=\"c-header__link\">About</a></li>" +
                      "<li class=\"c-header__item\" role=\"menuitem\"><a href=\"/blog\" class=\"c-header__link\">Blog</a></li>" +
                      "<li class=\"c-header__item\" role=\"menuitem\"><a href=\"/community\" class=\"c-header__link\">Community</a></li>" +
                      "<li class=\"c-header__item\" role=\"menuitem\"><a href=\"/faq\" class=\"c-header__link\">FAQ</a></li>" +
                 "</ul>" +
                "<a class=\"c-header__button\" href=\"/download\">Sign up</a>" +
               "</div>" +
              "</nav>" +
             "</div>" +
            "</header>"
}
