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

func GetHotBoardList() map[string]string {
	doc := GetDoc(HOT_BOARD_URL)

	// Parser
	// Get board name & href
	hot_board_list := make(map[string]string)
	doc.Find(".board").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		board_name := s.Find(".board-name").Text()
		hot_board_list[board_name] = href
	})
	for i, v := range hot_board_list {
		log.Debugf("%v: %v", i, v)
	}

	return hot_board_list
}
