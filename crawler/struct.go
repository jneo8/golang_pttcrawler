package crawler

import (
	"github.com/PuerkitoBio/goquery"
)

type Board struct {
	Urls     []string
	Name     string
	IndexUrl string
}

type Article struct {
	Doc     *goquery.Document
	ID      string
	RawHtml string
	Url     string
	Title   string
	Author  string
}
