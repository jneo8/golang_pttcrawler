package ptt

import (
    "net/http"
    "fmt"
    "log"
    // "io/ioutil"
    "golang.org/x/net/html"
    "reflect"
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
            switch tt {
                case html.ErrorToken:
                    fmt.Println("End")
                    break tokens_loop
                default:
                    tn, _ := z.TagName()
                    t := z.Token()
                    key, value, more_attr := z.TagAttr()
                    // TODO
                    if true {
                        fmt.Printf("---------\n")
                        fmt.Printf("tagename: %s data: %s\n", tn, t.Data)
                        fmt.Printf("key: %s\n", key)
                        fmt.Printf("value: %s type: %t\n", value, reflect.TypeOf(value))
                        fmt.Printf("more_attr: %t\n", more_attr)
                        fmt.Printf("---------\n")
                    }

                    // fmt.Printf("---------\n")
                    // fmt.Printf("tagename: %s data: %s\n", tn, t.Data)
                    // fmt.Printf("key: %s\n", key)
                    // fmt.Printf("value: %s type: %t\n", value, value)
                    // fmt.Printf("more_attr: %t\n", more_attr)
                    // fmt.Printf("---------\n")
            }
        }
    return resp
}


func GetBoardList() {
    url := "https://www.ptt.cc/bbs/hotboards.html"
    resp := GetResp(url)
    fmt.Printf("%T", resp)
}

