package crawler

import (
	mgo "github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

type pttRepository struct {
	Session *mgo.Session
}

func insertArticle(article *Article) {
	log.Debugf("Insert article: %v", article.ID)
}
