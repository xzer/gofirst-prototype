package web 

import (
    "html/template"
    "net/http"
    "fmt"
    "strconv"
    //"io/ioutil"
    "github.com/PuerkitoBio/goquery"
    "os"
    
)

type Data struct{
    Name string
    Age int
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func createHandler(targetFile string) handlerFunc{
    return func(w http.ResponseWriter, r *http.Request){
        //fmt.Fprintf(w, "request for %s", targetFile)
        params := r.URL.Query()
        
        filePath := "webroot" + targetFile
        
        data := Data{"", 0}
        data.Name = params["name"][0]
        data.Age, _ = strconv.Atoi(params["age"][0])
        
        t, _ := template.ParseFiles(filePath)
        t.Execute(w, &data)
    }
}

func createHandlerX(targetFile string) handlerFunc{
    return func(w http.ResponseWriter, r *http.Request){
        filePath := "webroot" + targetFile
        fmt.Printf("request for %s\n", filePath)
        reader, _ := os.Open(filePath)
        doc, _ := goquery.NewDocumentFromReader(reader)
        clonedDoc := doc.Clone()
        fmt.Fprintf(w, Render(clonedDoc))
    }
}

func StartDispatchter(addr string, rules *Rules){
    
    for _, rule := range rules.rules {
        if rule.url == "/abc" {
            http.HandleFunc(rule.url, createHandler(rule.target))
        }else{
            http.HandleFunc(rule.url, createHandlerX(rule.target))
        }
    }
    
    fmt.Printf("Started listening at %s\n", addr)
    var err = http.ListenAndServe(addr, nil)
    if err != nil {
        fmt.Printf("failed to start listenning due to \"%s\"", err)
    }
}

