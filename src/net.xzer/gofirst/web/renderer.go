package web

import (
    "github.com/PuerkitoBio/goquery"
    "golang.org/x/net/html"
)

type Renderer struct{
    funcs []RenderFunc
}

func (r *Renderer) Add(selector string, value string) {
    r.funcs = append(r.funcs, func(node *goquery.Selection){
        node.Find(selector).Each(func(i int, s *goquery.Selection){
            tn := html.Node{Type: html.TextNode, Data: value}
            s.Empty()
            s.AppendNodes(&tn)
        })
    })
}
