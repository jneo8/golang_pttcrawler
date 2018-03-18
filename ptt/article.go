package ptt


import (
    "github.com/PuerkitoBio/goquery"
    "github.com/fatih/color"
    "regexp"
    "strings"
)

type Comment struct {
    Category string  // pushing or boosting
    Author string
    DateTime string
    Content string
}

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
    Comments []*Comment
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
    // End ID

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

    // For get content later.
    header.Remove()
    header_right := doc.Find(".article-metaline-right")
    header_right.Remove()
    // End Header

    // Get IP
    ip := DEFAULT_IP
    re_ip, _ := regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+")
    doc.Find("#main-content").Find(".f2").Each(func(i int, s * goquery.Selection) {
        ip_text := ""
        if strings.Contains(s.Text(), "來自") || strings.Contains(s.Text(), "From") || strings.Contains(s.Text(), "編輯") {
            ip_text = re_ip.FindString(s.Text())
        }
        if (re_ip.FindString(ip_text) != "") {
            ip = ip_text
        }
    })
    article.IP = ip
    // End IP


    // Get Pushing && boosting
    pushing, boosting := GetComments(doc)
    article.Pushing = pushing
    article.Boosting = boosting

    // For get Content later.
    push := doc.Find(".push")
    push.Remove()
    // End Pushing && boosting

    // Get Content
    content := doc.Find("#main-content").Text()
    content = strings.Split(content, "※ 發信站:")[0]
    article.Content = content


    color.Green("%#v\n", article)
    return article
}


func GetComments(doc *goquery.Document) (int, int){
    push := doc.Find(".push")
    // Get Pushing && Boosting
    pushing := push.Find(".push-tag:contains('推 ')").Size()
    boosting := push.Find(".push-tag:contains('噓 ')").Size()

    // Get Coment
    // Comments := make([]*Comment)
    push.Each(func(i int, s * goquery.Selection) {
        comment := &Comment{}
        // Get Content
        content := s.Find(".push-content").Text()
        comment.Content = content
        // End Content


        // Get Category
        category := s.Find(".push-tag").Text()
        category = strings.Trim(category, ":")
        category = strings.TrimSpace(category)
        if (category == "推") {
            category = "pushing"
        } else if (category == "噓") {
            category = "boosting"
        } else {
            category = "no-emotion"
        }
        comment.Category = category
        // End category


        color.Magenta("%#v", comment)
    })
    color.Cyan("p: %d, b: %d", pushing, boosting)
    return pushing, boosting

}
