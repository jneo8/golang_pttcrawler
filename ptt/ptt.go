package ptt

import (
    "fmt"
    "log"
    "net/http"
    // "math"
    "strings"
    "strconv"
    "regexp"
    "github.com/PuerkitoBio/goquery"
)

const (
    BASE_URL = "https://www.ptt.cc/bbs/"
    ARTICLE_BASE_URL = "https://www.ptt.cc"
    HOT_BOARD_URL = "https://www.ptt.cc/bbs/hotboards.html"

    // default value
    DEFAULT_AUTHOR_NAME = "DEFAULT_AUTHOR"
    DEFAULT_TITLE = "DEFAULT_TITLE"
)

type Article struct {
    // ID       string
    Board    string
    Title    string
    Url      string
    // Content  string
    Author   string
    // DateTime string
    Nrec     int
    // doc      *goquery.Document
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

    return hot_board_list
}


func GetTitleList(board string) {
    url := BASE_URL + board + "/index.html"
    doc := GetDoc(url)
    articles := make([]*Article, 0)

    doc.Find(".r-ent").Each(func(i int, s *goquery.Selection) {
        article := &Article{Board: board}

        // Url
        url := s.Find(".title").Find("a")
        if len(url.Nodes) > 0 {
            href, _ := url.Attr("href")
            fmt.Printf("%s\n", href)
            GetArticle(ARTICLE_BASE_URL + href, board)
        }

        // Title
        title := strings.TrimSpace(s.Find(".title").Text())
        article.Title = title

        // nrec
        nrec := s.Find(".nrec")
        if len(nrec.Nodes) > 0 {
            nrec_str := nrec.Text()
            nrec_num, _ := strconv.Atoi(nrec_str)
            article.Nrec = nrec_num
        }
        // date
        author := s.Find(".author")
        if len(author.Nodes) > 0{
            article.Author = author.Text()
        }

        articles = append(articles, article)
    })

    // Get prePage & nextPage link
    prePage := doc.Find(".action-bar").Find("a:contains('‹ 上頁')")
    if len(prePage.Nodes) > 0 {
        href, _ := prePage.Attr("href")
        fmt.Printf(href)
    }

    nextPage := doc.Find(".action-bar").Find("a:contains('下頁 ›')")
    if len(nextPage.Nodes) > 0 {
        href, _ := nextPage.Attr("href")
        fmt.Printf(href)
    }

    for idx, v := range articles {
        fmt.Printf("%v %v\n", idx, v)
    }

}


func GetArticle(url string, board string) {
    // init article
    article := &Article{Board: board}
    doc := GetDoc(url)
    fmt.Printf("url: %s\n", url)

    // Author 
    author_origin := doc.Find(".article-metaline").Find(".article-meta-value").Eq(0).Text()
    // If not author, give default NAME
    if len(author_origin) == 0 {
        author_origin = DEFAULT_AUTHOR_NAME
    }
    re, _ := regexp.Compile("\\s\\([\\S\\s]+?\\)|\\s\\(\\)")
    // remove ()
    author := re.ReplaceAllString(author_origin, "")
    article.Author = author
    fmt.Printf("author: %s\n", author)

    title := doc.Find(".article-metaline").Find(".article-meta-value").Eq(1).Text()
    if len(title) == 0 {
        title = DEFAULT_TITLE
    }
    article.Title = title
    fmt.Printf("title: %s\n", title)
    // article
    fmt.Printf("article: %v\n----\n", article)
}

