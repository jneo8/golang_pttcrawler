package crawler

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
	"sync"
)

func Init() {
	boards := GetHotBoardList()
	boardChan := make(chan Board, 10)
	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go crawlerWorker(&wg, i, boardChan)
	}

	for _, board := range boards {
		boardChan <- board
	}
	// stop give job.
	close(boardChan)
	wg.Wait()
}

func crawlerWorker(wg *sync.WaitGroup, workerID int, boardChan <-chan Board) {
	defer wg.Done()
	session, err := createSession("localhost:27017")
	if err != nil {
		log.Panic(err)
	}
	pttRepository := PttRepository{Session: session}
	defer pttRepository.close()

	for board := range boardChan {
		urlChan := make(chan string)
		articleChan := make(chan Article)
		jobWg := sync.WaitGroup{}
		jobWg.Add(2)
		log.Debugf("worker: %v, %v", workerID, board)
		go board.getUrls(&jobWg, urlChan)
		go board.getArticle(&jobWg, urlChan, articleChan)

		for article := range articleChan {
			pttRepository.insertArticle(&article)
		}
		jobWg.Wait()
	}
}

func (board *Board) getArticle(wg *sync.WaitGroup, urlChan <-chan string, articleChan chan<- Article) {
	defer wg.Done()

	for url := range urlChan {
		article := Article{}

		// Article Url
		article.Url = url

		// Article ID
		idCompile, _ := regexp.Compile("bbs/(.*).html$")
		articleID := idCompile.FindString(url)
		articleID = strings.Trim(articleID, "bbs/")
		articleID = strings.Trim(articleID, ".html")
		article.ID = articleID

		// Get Doc
		doc := GetDoc(url)
		// article.Doc = doc

		// Article Raw HTML
		rawHtml, err := doc.Html()
		if err != nil {
			log.Error(err)
			article.RawHtml = ""
		} else {
			article.RawHtml = rawHtml
		}

		title, author := getInfo(doc)
		article.Title = title
		article.Author = author

		articleChan <- article

		// TODO parser html
	}
	close(articleChan)
}

func (board *Board) getUrls(wg *sync.WaitGroup, urlChan chan<- string) {
	defer wg.Done()
	log.Debugf("%v", board)

	doc := GetDoc(board.IndexUrl)

	// For now it is for test, for production we will check if data in DB.
	for idx := 1; idx <= 1; idx++ {
		// Get article url in page.
		doc.Find(".r-ent").Each(
			func(i int, s *goquery.Selection) {
				a := s.Find(".title").Find("a")
				if len(a.Nodes) > 0 {
					articleHref, _ := a.Attr("href")
					articleHref = ARTICLE_BASE_URL + articleHref
					urlChan <- articleHref
				}
			},
		)
		// NextPage
		nextPage := doc.Find(".action-bar").Find("a:contains('‹ 上頁')")
		if len(nextPage.Nodes) > 0 {
			nextPageHref, _ := nextPage.Attr("href")
			nextPageHref = BASE_URL + nextPageHref
			doc = GetDoc(nextPageHref)
		} else {
			log.Warning("NextPage not find %v:%v", board.Name, idx)
			break
		}
	}
	close(urlChan)
}

// Get info from article.
func getInfo(doc *goquery.Document) (string, string) {

	// Header, Get author, title, date
	header := doc.Find(".article-metaline")

	// Get title
	title := header.Find(".article-meta-value").Eq(1).Text()
	author := header.Find(".article-meta-value").Eq(0).Text()
	re_author, _ := regexp.Compile("\\s\\([\\S\\s]+?\\)|\\s\\(\\)")
	author = re_author.ReplaceAllString(author, "")
	log.Debugf("Title: %v", title)
	log.Debugf("Author: %v", author)
	return title, author
}
