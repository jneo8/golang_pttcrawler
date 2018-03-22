package ptt


import (
    "github.com/PuerkitoBio/goquery"
    "github.com/fatih/color"
    "regexp"
    "strings"
    "time"
)

type Comment struct {
    Tag string  // pushing or boosting or no-emotion
    Author string
    DateTime string
    Content string
    CrawlerTime time.Time
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

func GetArticles(fishes []*Fish) []*Article{
    articles := make([]*Article, 0)
    for index := range fishes {
        color.Yellow("Get article: %s", fishes[index].Url)
        article := GetArticle(fishes[index].Url)
        articles = append(articles, article)
    }
    return articles
}

func GetArticle(url string) *Article {
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


    // Push

    push := doc.Find(".push")

    // Get Pushing && Boosting
    pushing := push.Find(".push-tag:contains('推 ')").Size()
    boosting := push.Find(".push-tag:contains('噓 ')").Size()
    article.Pushing = pushing
    article.Boosting = boosting
    // End Pushing && boosting

    // Get Comments
    comments := GetComments(doc)
    article.Comments = comments
    // End Comments

    // For get Content later.
    push.Remove()

    // End Push

    // Get Content
    content := doc.Find("#main-content").Text()
    content = strings.Split(content, "※ 發信站:")[0]
    article.Content = content

    return article
}


func GetComments(doc *goquery.Document) []*Comment {
    push := doc.Find(".push")

    // Get Coment
    comments := make([]*Comment, 0)
    push.Each(func(i int, s * goquery.Selection) {
        comment := &Comment{}
        // Get Content
        content := s.Find(".push-content").Text()
        comment.Content = content
        // End Content

        // Get Tag
        tag := s.Find(".push-tag").Text()
        tag = strings.Trim(tag, ":")
        tag = strings.TrimSpace(tag)
        if (tag == "推") {
            tag = "pushing"
        } else if (tag == "噓") {
            tag = "boosting"
        } else {
            tag = "no-emotion"
        }
        comment.Tag = tag
        // End Tag

        // Author
        author := s.Find(".push-userid").Text()
        comment.Author = author
        // End Author

        // Datetime 
        datetime := s.Find(".push-ipdatetime").Text()
        datetime = strings.TrimSpace(datetime)
        comment.DateTime = datetime
        // End Datetime

        // Get CrawlerTime
        crawlertime := time.Now()
        comment.CrawlerTime = crawlertime
        // End CrawlerTime

        comments = append(comments, comment)
    })
    return comments

}
