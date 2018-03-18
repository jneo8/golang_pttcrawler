package ptt


import (
    // "github.com/PuerkitoBio/goquery"
    "github.com/fatih/color"
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
        color.Green("Get article: %s", fish.Board.Urls[index])
        GetArticle(fish.Board.Urls[index])
    }
}

func GetArticle(url string) {
    color.Green("%s", url)

    // Get Doc
    doc := GetDoc(url)

    article := &Article{}
    article.Url = url

    // Get Author
    author_origin := doc.Find(".article-metaline").Find(".article-meta-value").Eq(0).Text()
    if len(author_origin) == 0 {
        author_origin = DEFAULT_AUTHOR_NAME
    }

    color.Green("%#v\n", article)
}
