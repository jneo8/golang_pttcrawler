package ptt

import (
    "fmt"
    "log"
    "net/http"
    "github.com/PuerkitoBio/goquery"
)

const (
    BASE_URL = "https://www.ptt.cc/bbs/"
)

func GetResp(url string) (*goquery.Document) {
    // Add cookie
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal(err)
    }

    cookie := http.Cookie{
        Name:  "over18",
        Value: "1",
    }
    req.AddCookie(&cookie)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    if resp.StatusCode != http.StatusOK {
        log.Fatal(resp.StatusCode)
    }

    doc, err := goquery.NewDocumentFromResponse(resp)
    if err != nil {
        log.Fatal(err)
    }

    return doc
}

func GetBoardList() {
    url := "https://www.ptt.cc/bbs/hotboards.html"
    resp := GetResp(url)
    fmt.Printf("%T", resp)
}

