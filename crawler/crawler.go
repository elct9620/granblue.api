package crawler

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// Fetcher
type Parser interface {
}

// Crawler
type Crawler struct {
	Source   string
	Pagename string

	parser *Parser
	body   *goquery.Selection
}

// TODO(elct9620): Implement crawler which can scan gbf-wiki data
func New() *Crawler {
	return &Crawler{Source: SOURCE_URL}
}

// Set parser to handle fetched html page
func (c *Crawler) SetParser(parser *Parser) {
	c.parser = parser
}

// Set which page the fetcher should fetch html
func (c *Crawler) SetPage(page string) {
	c.Pagename = page
}

// Fetch the page from source
func (c *Crawler) Fetch() error {

	res, err := http.Get(fmt.Sprintf("%s?%s", c.Source, c.Pagename))
	if err != nil {
		return err
	}

	defer res.Body.Close()
	readInUTF8 := transform.NewReader(res.Body, japanese.EUCJP.NewDecoder())

	doc, err := goquery.NewDocumentFromReader(readInUTF8)
	if err != nil {
		return err
	}

	c.body = doc.Find("#body")

	return nil
}
