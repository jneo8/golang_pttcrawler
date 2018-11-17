package crawler

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetDoc(url string) *goquery.Document {
	// Add cookie
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	cookie := http.Cookie{
		Name:  "over18",
		Value: "1",
	}
	req.AddCookie(&cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func GetHotBoardList() []Board {
	doc := GetDoc(HOT_BOARD_URL)

	// Parser
	// Get board name & href
	hotBoardList := []Board{}
	doc.Find(".board").Each(func(i int, s *goquery.Selection) {
		href, err := s.Attr("href")
		if err != nil {
			logrus.Panic(err)
		}
		boardName := s.Find(".board-name").Text()
		hotBoardList = append(hotBoardList, Board{IndexUrl: BASE_URL + href, Name: boardName})

	})
	for i, v := range hotBoardList {
		log.Debugf("%v: %v", i, v)
	}

	return hotBoardList
}
