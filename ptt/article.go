package ptt


import (
    "github.com/PuerkitoBio/goquery"
    "github.com/fatih/color"
    "regexp"
    "strings"
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

func GetArticles(fish *Fish) {
    for index := range fish.Board.Urls {
        color.Yellow("Get article: %s", fish.Board.Urls[index])
        GetArticle(fish.Board.Urls[index])
    }
}

func GetArticle(url string) {
    color.Green("%s", url)

    // Get Doc
    doc := GetDoc(url)

    article := &Article{}
    article.Url = url


    // Header
    // Get author, title, date in Header
    header := doc.Find(".article-metaline")

    // Get Title
    title := header.Find(".article-meta-value").Eq(1).Text()
    if len(title) == 0 {
        title = DEFAULT_TITLE
    }
    article.Title = title
    // End Title

    // Get Author
    origin_author := header.Find(".article-meta-value").Eq(0).Text()
    if len(origin_author) == 0 {
        origin_author = DEFAULT_AUTHOR_NAME
    }
    // Remove () in origin_author
    re, _ := regexp.Compile("\\s\\([\\S\\s]+?\\)|\\s\\(\\)")
    author := re.ReplaceAllString(origin_author, "")
    article.Author = author
    // End Author

    // Get Datetime
    datetime := header.Find(".article-meta-value").Eq(2).Text()
    article.DateTime = datetime
    // End Datetime

    header.Remove()

    // Get ip
    ip := DEFAULT_IP
    doc.Find("#main-content").Find(".f2").Each(func(i int, s * goquery.Selection) {
        if strings.Contains(s.Text(), "來自") {
            ip = strings.Split(s.Text(), "來自: ")[1]
        }
    })
    color.Cyan("%s", ip)



    color.Green("%#v\n", article)
}
