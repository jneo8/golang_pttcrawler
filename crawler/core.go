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
		log.Debugf("worker: %v, %v", workerID, board)
	}
}

func getBoard(wg *sync.WaitGroup, board Board) {
	log.Debugf("%v", board)
}
