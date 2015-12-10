package web

import (
    //"fmt"
)

type Rule struct {
    url string
    target string
}

type Rules struct {
    rules []Rule
}

func (r *Rules) Add(url string, target string) {
    r.rules = append(r.rules, Rule{url, target})
}


