package web

import (
    "github.com/PuerkitoBio/goquery"
    //"golang.org/x/net/html"
    "fmt"
    "reflect"
)

type RenderFunc func(node *goquery.Selection)

var snippetMap map[string]reflect.Type

func init(){
    snippetMap = make(map[string]reflect.Type)
}

func RegisterSnippet(name string, snippet interface{}){
    snippetMap[name] = reflect.TypeOf(snippet)
}

func Render(doc *goquery.Selection) string{
    doc.Find("[gf-snippet]").Each(func(i int, s *goquery.Selection){
        fmt.Println(s.Html())
        snippetName, _ := s.Attr("gf-snippet")
        typ := snippetMap[snippetName]
        instanceValue := reflect.New(typ)
        rm := instanceValue.MethodByName("Render")
        rtn := rm.Call([]reflect.Value{})
        renderFuncValue := rtn[0]
        renderer := renderFuncValue.Interface().(Renderer)
        for _, fn := range renderer.funcs {
            fn(s)
        }
    })
    html, _ := doc.Html()
    return html
}
