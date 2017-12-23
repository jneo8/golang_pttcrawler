package ptt

import (
    "net/http"
    "fmt"
    "io/ioutil"
)

const (
    BASE_URL = "https://www.ptt.cc/bbs/"
)

func GetDoc(url string) ([]byte, error) {
    // Add cookie
    req, err := http.NewRequest("GET", url, nil)
    cookie := http.Cookie {
        Name: "over18",
        Value: "1",
    }
    req.AddCookie(&cookie)

    // Send req
    resp, err := http.DefaultClient.Do(req)

    // Read Body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    resp.Body.Close()
    return body, err
}


func GetArticles(board string) {
    idx := "/index.html"
    url := BASE_URL + board + idx
    fmt.Println(url)
    body, err := GetDoc(url)
    fmt.Printf("%s", body)
    fmt.Printf("%s", err)
}

func GetBoardList() {
    url := "https://www.ptt.cc/cls/1"
    body, err := GetDoc(url)
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Printf("%s", body)
    // Todo
    // parser
}
