package main

import "fmt"
import "net/http"
import "io/ioutil"

func main() {
    url := "http://news.sdtv.com.tw/News_menu2.html?today=2009-03-04"
    fmt.Println(getContentWithURL(url))
}

func getContentWithURL(url string) (string){
    resp, err := http.Get("http://news.sdtv.com.tw/News_menu2.html?today=2009-03-04")
    if err != nil {
    }

    defer resp.Body.Close()
    bytes, err := ioutil.ReadAll(resp.Body)
    return string(bytes)
}
