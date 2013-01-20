package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    "log"
    "html"
    //"sdtv"
)

func main() {
    path := "../data/News_detail.html"
    str, err := ioutil.ReadFile(path)
    cherr(err)

    ic, err := iconv.Open("UTF-8", "BIG-5")
    defer ic.Close()

    newstr, err := ic.Conv(string(str))
    cherr(err)
    fmt.Printf("new: %s\n", newstr)
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

func cherr(err error) {
    if err != nil {
	panic(err)
    }
}
