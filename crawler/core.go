package crawler

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
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
	for board := range boardChan {
		urlChan := make(chan string)
		jobWg := sync.WaitGroup{}
		jobWg.Add(2)
		log.Debugf("worker: %v, %v", workerID, board)
		go board.getUrls(&jobWg, urlChan)
		go board.getArticle(&jobWg, urlChan)
		jobWg.Wait()
	}
}

func (board *Board) getArticle(wg *sync.WaitGroup, urlChan <-chan string) {
	defer wg.Done()
	for url := range urlChan {
		log.Debugf("article: %v", url)
	}
}

func (board *Board) getUrls(wg *sync.WaitGroup, urlChan chan<- string) {
	defer wg.Done()
	log.Debugf("%v", board)

	doc := GetDoc(board.IndexUrl)

	// For now it is for test, for production we will check if data in DB.
	for idx := 1; idx <= 10; idx++ {
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
			log.Infof("NextPageHref: %v", nextPageHref)
			doc = GetDoc(nextPageHref)
		} else {
			log.Warning("NextPage not find %v:%v", board.Name, idx)
			break
		}
	}
	close(urlChan)
}