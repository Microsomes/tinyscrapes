package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	)

	type UANews struct {
		Title  string
		Date   string
		Source string
		Icon   string
	}

	// On every a element which has href attribute call callback
	c.OnHTML("body", func(e *colly.HTMLElement) {
		var Articles []UANews
		e.ForEach(".event", func(i int, e *colly.HTMLElement) {
			var title = e.DOM.Find(".title")
			e.ForEach("a", func(i int, e *colly.HTMLElement) {
				fmt.Print(e.DOM.Attr("href"))
			})

			Articles = append(Articles, UANews{
				Title: title.Text(),
			})
		})

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://liveuamap.com")
}
