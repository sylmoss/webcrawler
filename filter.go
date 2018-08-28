package main

import "regexp"

type Filter struct {
    regex   string
}

func NewFilter(regex string) *Filter {
    return &Filter{
        regex: regex,
    }
}

func (f *Filter) Find(body string) [][]string {
    return regexp.MustCompile(f.regex).FindAllStringSubmatch(body, -1)
}