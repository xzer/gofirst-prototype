package main 

import (
    "net.xzer/gofirst/web"
)

type MySnippet struct{
    x int
    y int
}

func init(){
    web.RegisterSnippet("MySnippet", MySnippet{})
}

func (s *MySnippet) Render() web.Renderer{
    renderer := web.Renderer{}
    renderer.Add(".x-name", "xzer")
    renderer.Add(".x-age", "33")
    return renderer
}

func main() {
    web.StartDispatchter(":8120", createRules())
}

func createRules() *web.Rules{
    rules := web.Rules{}
    rules.Add("/abc", "/abc.html")
    rules.Add("/xyz", "/xyz.html")
    return &rules
}

