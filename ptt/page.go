package ptt

import (
    "github.com/PuerkitoBio/goquery"
    "github.com/fatih/color"
)


type Page struct {
    Index int
    NextUrl  string
    Url   string
    Doc   *goquery.Document
}


func GetPages(board string, index int, url string, max int) []*Page{
    // Gen url.
    if (index == 0) {
        url = BASE_URL + "/bbs/" + board + "/index.html"
        max -= 1
    } else {
        url = BASE_URL + url
    }

    color.Blue(" Start get: %s-%d/%d: %s\n", board, index, max, url)
    pages := make([]*Page, 0)
    doc := GetDoc(url)

    if (index < max) {
        // Get nextPage
        nextPage := doc.Find(".action-bar").Find("a:contains('‹ 上頁')")

        if len(nextPage.Nodes) > 0 {
            href, _ := nextPage.Attr("href")
            index += 1
            now_page := &Page{Index: index, Url: url, NextUrl: href, Doc: doc}

            next_pages := GetPages(board, index, href, max)
            pages = append([]*Page{now_page}, next_pages...)
        } else {
            now_page := &Page{Index: index, Url: url, Doc: doc}
            pages = append(pages, now_page)
        }
    } else {

        now_page := &Page{Index: index, Url: url, Doc: doc}
        pages = append(pages, now_page)
    }
    return pages

}
