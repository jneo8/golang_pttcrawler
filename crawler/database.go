package crawler

import (
	log "github.com/sirupsen/logrus"
)

func insertArticle(article *Article) {
	log.Debugf("Insert article: %v", article.ID)
}
