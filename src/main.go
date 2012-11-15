package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    "log"
    "html"
    "sdtv"
)

func main() {
    path := "../data/News_detail.html"
    str, err := ioutil.ReadFile(path)
    if err != nil {
	log.Fatal("Can not open file %s", path)
    }

    detail := sdtv.Detail{}
    d := detail.Parse([]byte(str))
    fmt.Printf("Title: %s\nDescription: %s\nDate: %s\nWMV: %s\n", d.Title, d.Description, d.Date, d.WMV)
}

func foo() {
    /*
    url := "http://news.sdtv.com.tw/News_menu2.html?today=2009-03-04"
    respStr := getContentWithURL(url)
    */
    path := "../data/News_menu2.html"
    respStr, err := ioutil.ReadFile(path)
    if err != nil {
	log.Fatal("Can not open file, %s", path)
	return
    }
    r, _ := regexp.Compile("News_detail.html\\?sn=\\d+")
    for _, e := range r.FindAll([]byte(respStr), len(respStr)) {
	fmt.Printf("%s\n", html.UnescapeString(string(e)))
    }
}

func getContentWithURL(url string) (string){
    resp, err := http.Get("http://news.sdtv.com.tw/News_menu2.html?today=2009-03-04")
    if err != nil {
    }

    defer resp.Body.Close()
    bytes, err := ioutil.ReadAll(resp.Body)
    return string(bytes)
}
