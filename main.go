package main

import "fmt"
import "net/http"
import "io/ioutil"
import "unicode/utf8"

func main() {
    url := "http://news.sdtv.com.tw/News_menu2.html?today=2009-03-04"
    respStr := getContentWithURL(url)
    /* 
    FIXME:
	Convert text encoding with libiconv.
	how to call libiconv functions in Go?, http://goo.gl/92tid
    */

}

func getContentWithURL(url string) (string){
    resp, err := http.Get("http://news.sdtv.com.tw/News_menu2.html?today=2009-03-04")
    if err != nil {
    }

    defer resp.Body.Close()
    bytes, err := ioutil.ReadAll(resp.Body)
    return string(bytes)
}
