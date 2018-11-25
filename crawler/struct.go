package crawler

type Board struct {
	Urls     []string
	Name     string
	IndexUrl string
}

type Article struct {
	ID      string
	RawHtml string
	Url     string
}
