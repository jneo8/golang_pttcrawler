package crawler

import (
	mgo "github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

type PttRepository struct {
	Session *mgo.Session
}

func createSession(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	return session, nil
}

func (repo *PttRepository) close() {
	repo.Session.Close()
}

func (repo *PttRepository) insertArticle(article *Article) {
	log.Debugf("Insert article: %v", article.ID)
	repo.collection().Insert(article)
}

func (repo *PttRepository) collection() *mgo.Collection {
	return repo.Session.DB("ptt").C("articles")
}
