package ptt

import (
    // "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/fatih/color"
)


type Page struct {
    Index int
    NextUrl  string
    Url   string
}

type PageList struct {
    Pages []*Page
    Board string
}

func GetHotBoardList() map[string]string {
    doc := GetDoc(HOT_BOARD_URL)

    // Parser
    // Get board name & href
    hot_board_list := make(map[string]string)
    doc.Find(".board").Each(func(i int, s *goquery.Selection) {
        href, _ := s .Attr("href")
        board_name := s.Find(".board-name").Text()
        hot_board_list[board_name] = href
    })

    return hot_board_list
}

func GetAllDocUrl(board string, index int, url string, max int) {
    pages := GetAllPageUrl(board, index, url, max)

    for index, page := range pages {
        color.Red("%d: %s", index, page.Url)
    }
    pagelist := PageList{Board: board, Pages: pages}
    color.Cyan("%s", pagelist)

}

func GetAllPageUrl(board string, index int, url string, max int) []*Page{
    // Gen url.
    if (index == 0) {
        url = BASE_URL + "/bbs/" + board + "/index.html"
        max -= 1
    } else {
        url = BASE_URL + url
    }

    color.Blue("%s-%d/%d: %s\n", board, index, max, url)
    pages := make([]*Page, 0)

    if (index < max) {
        doc := GetDoc(url)

        // Get nextPage
        nextPage := doc.Find(".action-bar").Find("a:contains('‹ 上頁')")

        if len(nextPage.Nodes) > 0 {
            href, _ := nextPage.Attr("href")
            index += 1
            now_page := &Page{Index: index, Url: url, NextUrl: href}

            next_pages := GetAllPageUrl(board, index, href, max)
            pages = append([]*Page{now_page}, next_pages...)
        } else {
            now_page := &Page{Index: index, Url: url, NextUrl: ""}
            pages = append(pages, now_page)
        }
    } else {
        now_page := &Page{Index: index, Url: url, NextUrl: ""}
        pages = append(pages, now_page)
    }
    return pages

}
