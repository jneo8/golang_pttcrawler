package ptt

import (
    "net/http"
    "fmt"
    "log"
    // "io/ioutil"
    "golang.org/x/net/html"
    // "strings"
    // "reflect"
)

const (
    BASE_URL = "https://www.ptt.cc/bbs/"
)

func GetResp(url string) (*http.Response) {
    // Add cookie
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal(err)
    }
    // cookie := http.Cookie {
    //     Name: "over18",
    //     Value: "1",
    // }
    // req.AddCookie(&cookie)

    // Send req
    resp, err := http.DefaultClient.Do(req)

    // Read Body
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(resp.Body)
    // bytes, err := ioutil.ReadAll(resp.Body)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // text := string(bytes)
    // fmt.Printf("%T\n", text)
    defer resp.Body.Close()
    fmt.Println(resp.Body)

    z := html.NewTokenizer(resp.Body)
    tokens_loop:
        for {
            tt := z.Next()
            fmt.Printf("%T\n", z)
            switch tt {
                case html.ErrorToken:
                    fmt.Println("End")
                    break tokens_loop
                default:
                    tn, _ := z.TagName()
                    t := z.Token()
                    fmt.Printf("tagename: %s data: %s\n", tn, t.Data)
            }
        }
    return resp
}


func GetBoardList() {
    url := "https://www.ptt.cc/bbs/hotboards.html"
    resp := GetResp(url)
    fmt.Printf("%T", resp)
}

