package ptt

import (
    "github.com/PuerkitoBio/goquery"
    "github.com/fatih/color"
)

type Board struct {
    Urls []string
}

func GetBoard(board_name string, max int) Board{
    // Get page's doc.
    pages := GetPages(board_name, 0, "", max)

    urls := make([]string, 0)

    // Get urls in every page
    for index, page := range pages {
        color.Red("%d: Get url in %s\n", index, page.Url)
        if (page.Doc != nil) {
            page.Doc.Find(".r-ent").Each(func(i int, s *goquery.Selection) {
                a := s.Find(".title").Find("a")
                if len(a.Nodes) > 0 {
                    href, _ := a.Attr("href")
                    href = ARTICLE_BASE_URL + href
                    urls = append(urls, href)
                }
            })
        }
    }
    board := Board{Urls: urls}
    return board
}
