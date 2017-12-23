package ptt

import (
    "net/http"
    "fmt"
    "log"
    // "io/ioutil"
    "golang.org/x/net/html"
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
    // bytes, err := ioutil.ReadAll(resp.Body)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // text := string(bytes)
    // fmt.Println(text)

    fmt.Printf("%r\n\n\n", resp)
    defer resp.Body.Close()
    return resp
}


func GetBoardList() {
    url := "https://www.ptt.cc/bbs/hotboards.html"
    resp := GetResp(url)
    htmlTokens := html.NewTokenizer(resp.Body)
    token_loop:
        for {
            tokenType := htmlTokens.Next()
            fmt.Printf("tokenType: %s\n", tokenType)
            switch tokenType {
                case html.ErrorToken:
                    fmt.Println("End")
                    break token_loop
                case html.TextToken:
                    fmt.Println(tokenType)
                default:
                    fmt.Println(tokenType)
            }
        }
    // Todo
    // parser
}

