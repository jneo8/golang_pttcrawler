package ptt


import (
    "github.com/PuerkitoBio/goquery"
    "github.com/fatih/color"
    "regexp"
    "strings"
)

type Article struct {
    ID       string
    Title    string
    Url      string
    Content  string
    Author   string
    DateTime string
    Pushing  int
    Boosting int
    IP       string
}

func GetArticles(fish *Fish) {
    for index := range fish.Board.Urls {
        color.Yellow("Get article: %s", fish.Board.Urls[index])
        GetArticle(fish.Board.Urls[index])
    }
}

func GetArticle(url string) *Article {
    color.Green("%s", url)

    article := &Article{}

    // Get ID
    re_id, _ := regexp.Compile("bbs/(.*).html$")
    id := re_id.FindString(url)
    id = strings.Trim(id, "bbs/")
    id = strings.Trim(id, ".html")
    article.ID = id

    // Get Doc
    doc := GetDoc(url)

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
    re_author, _ := regexp.Compile("\\s\\([\\S\\s]+?\\)|\\s\\(\\)")
    author := re_author.ReplaceAllString(origin_author, "")
    article.Author = author
    // End Author

    // Get Datetime
    datetime := header.Find(".article-meta-value").Eq(2).Text()
    article.DateTime = datetime
    // End Datetime

    header.Remove()
    header_right := doc.Find(".article-metaline-right")
    header_right.Remove()

    // Get IP
    ip := DEFAULT_IP
    re_ip, _ := regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+")
    doc.Find("#main-content").Find(".f2").Each(func(i int, s * goquery.Selection) {
        ip_text := ""
        if strings.Contains(s.Text(), "來自") || strings.Contains(s.Text(), "From") || strings.Contains(s.Text(), "編輯") {
            ip_text = re.FindString(s.Text())
        }
        if (re_ip.FindString(ip_text) != "") {
            ip = ip_text
        }
    })
    article.IP = ip
    // End IP


    // Get Pushing && boosting
    push := doc.Find(".push")
    pushing := push.Find(".push-tag:contains('推 ')").Size()
    boosting := push.Find(".push-tag:contains('噓 ')").Size()
    article.Pushing = pushing
    article.Boosting =boosting
    push.Remove()
    // End Pushing && boosting

    content := doc.Find("#main-content").Text()
    content = strings.Split(content, "※ 發信站:")[0]
    article.Content = content


    color.Green("%#v\n", article)
    return article
}
