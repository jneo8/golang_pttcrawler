package crawler

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

func Init() {
	boards := GetHotBoardList()
	boardChan := make(chan Board, 10)
	wg := sync.WaitGroup{}
	for i := 1; i <= 2; i++ {
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
		jobWg := sync.WaitGroup{}
		jobWg.Add(1)
		log.Debugf("worker: %v, %v", workerID, board)
		getBoard(&jobWg, board)
		jobWg.Wait()
	}
}

func getBoard(wg *sync.WaitGroup, board Board) {
	defer wg.Done()
	log.Debugf("%v", board)

	doc := GetDoc(board.IndexUrl)
	for i := 1; i <= 10; i++ {
		nextPage := doc.Find(".action-bar").Find("a:contains('‹ 上頁')")
		if len(nextPage.Nodes) > 0 {
			nextPageHref, _ := nextPage.Attr("href")
			nextPageHref = BASE_URL + nextPageHref
			log.Debugf("NextPageHref: %v", nextPageHref)
			doc = GetDoc(nextPageHref)
		} else {
			log.Warning("NextPage not find %v:%v", board.Name, i)
			break
		}
	}
}
