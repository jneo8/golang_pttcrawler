package ptt

import (
    "fmt"
    "log"
    "net/http"
    // "math"
    "strings"
    "strconv"
    "github.com/PuerkitoBio/goquery"
)

const (
    BASE_URL = "https://www.ptt.cc/bbs/"
    HOT_BOARD_URL = "https://www.ptt.cc/bbs/hotboards.html"
)

type Article struct {
    ID       string
    Board    string
    Title    string
    Content  string
    Author   string
    DateTime string
    Nrec     int
    Url      string
    doc      *goquery.Document
}

func GetDoc(url string) (*goquery.Document) {
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


func GetHotBoardList() map[string]string {
    doc := GetDoc(HOT_BOARD_URL)

    // Parser
    // Get board name & href
    hot_board_list := make(map[string]string)
    doc.Find(".board").Each(func(i int, s *goquery.Selection) {
        href, _ := s.Attr("href")
        board_name := s.Find(".board-name").Text()
        hot_board_list[board_name] = href
    })

    // Print
    for name, href := range hot_board_list {
        fmt.Printf("%s: %s\n", name, href)
    }

    return hot_board_list
}


func GetArticles(board string) {
    url := BASE_URL + board + "/index.html"
    doc := GetDoc(url)
    // articles := make([]*Article, 0)

    doc.Find(".r-ent").Each(func(i int, s *goquery.Selection) {
        // article := &Article{Board: board}

        // Title
        title := strings.TrimSpace(s.Find(".title").Text())

        // nrec
        nrec := s.Find(".nrec")
        if len(nrec.Nodes) > 0 {
            nrec_str := nrec.Text()
            nrec_num, _ := strconv.Atoi(nrec_str)
            fmt.Printf("%d : %s\n", nrec_num, title)
        }

    })

}

