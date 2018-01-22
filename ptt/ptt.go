package ptt

import (
    "fmt"
    "log"
    "net/http"
    // "math"
    // "strings"
    // "strconv"
    "regexp"
    "strings"
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
    Content  string
    Author   string
    DateTime string
    Pushing  int
    Boosting int
    IP       string
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
        // Url
        url := s.Find(".title").Find("a")
        if len(url.Nodes) > 0 {
            href, _ := url.Attr("href")
            fmt.Printf("%s\n", href)
            article := GetArticle(ARTICLE_BASE_URL + href, board)
            articles = append(articles, article)
        }
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
        fmt.Printf("%v %v %v\n", idx, v.Pushing, v.Boosting)
    }

}


func GetArticle(url string, board string) *Article {
    // init article
    article := &Article{Board: board}

    // Get doc
    doc := GetDoc(url)

    // Url
    article.Url = url
    fmt.Printf("url: %s\n", url)

    // Author 
    author_origin := doc.Find(".article-metaline").Find(".article-meta-value").Eq(0).Text()
    // If not author, give default NAME
    if len(author_origin) == 0 {
        author_origin = DEFAULT_AUTHOR_NAME
    }
    // remove () in author
    re, _ := regexp.Compile("\\s\\([\\S\\s]+?\\)|\\s\\(\\)")
    author := re.ReplaceAllString(author_origin, "")
    article.Author = author
    fmt.Printf("author: %s\n", author)

    // Title
    title := doc.Find(".article-metaline").Find(".article-meta-value").Eq(1).Text()
    if len(title) == 0 {
        title = DEFAULT_TITLE
    }
    article.Title = title
    fmt.Printf("title: %s\n", title)

    // Date
    datetime := doc.Find(".article-metaline").Find(".article-meta-value").Eq(2).Text()
    article.DateTime = datetime
    fmt.Printf("date: %s\n", datetime)

    // pushing & boosting
    push := doc.Find(".push")
    fmt.Printf("%v\n", push.Size())
    pushing := push.Find(".push-tag:contains('推 ')").Size()
    boosting := push.Find(".push-tag:contains('噓 ')").Size()
    article.Pushing = pushing
    article.Boosting = boosting

    // ip
    ip := "noip"
    doc.Find("#main-content").Find(".f2").Each(func(i int, s *goquery.Selection) {
        if strings.Contains(s.Text(), "來自") {
            ip = strings.Split(s.Text(), "來自: ")[1]
        }
    })
    article.IP = ip
    fmt.Printf("ip:%s", ip)


    // content
    header := doc.Find(".article-metaline")
    header.Remove()
    headerRight := doc.Find(".article-metaline-right")
    headerRight.Remove()
    pushs := doc.Find(".push")
    pushs.Remove()
    content := doc.Find("#main-content").Text()
    article.Content = content
    // fmt.Printf("content: %s\n", content)

    // article
    // fmt.Printf("article: %v\n----\n", article)
    return article
}

