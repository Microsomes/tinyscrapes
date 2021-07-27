package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"unicode/utf8"

	"github.com/gocolly/colly"
)

type Prayer struct {
	Month   string
	Day     string
	Sunrise string
	Fajr    string
	Zuhr    string
	Asr     string
	Maghrib string
	Isha    string
}

func CrawlBCM(cc chan []Prayer, monthRequested int) {
	var Prayers []Prayer
	var url = fmt.Sprintf("https://centralmosque.org.uk/wp-admin/admin-ajax.php?action=get_monthly_timetable&month=%d", monthRequested)
	fmt.Println(url)
	c := colly.NewCollector()
	c.SetRequestTimeout(100000000000)
	//request timeout set to 100 seconds

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, f *colly.HTMLElement) {
			var month = ""
			var day = ""
			var fajr = ""
			var sunrise = ""
			var zuhr = ""
			var asr = ""
			var maghrib = ""
			var isha = ""

			f.ForEach("td", func(i int, h *colly.HTMLElement) {
				if i == 0 {
					month = h.Text
				}
				if i == 1 {
					day = h.Text
				}
				if i == 2 {
					fajr = h.Text
				}
				if i == 4 {
					sunrise = h.Text
				}
				if i == 5 {
					zuhr = h.Text
				}
				if i == 7 {
					asr = h.Text
				}
				if i == 9 {
					maghrib = h.Text
				}
				if i == 11 {
					isha = h.Text
				}
			})

			if utf8.RuneCountInString(month) >= 4 {
				Prayers = append(Prayers, Prayer{
					Month:   month,
					Day:     day,
					Fajr:    fajr,
					Sunrise: sunrise,
					Zuhr:    zuhr,
					Asr:     asr,
					Maghrib: maghrib,
					Isha:    isha,
				})
			}
		})
		cc <- Prayers

	})
	c.Visit(url)
}

func processNAMAZ(w http.ResponseWriter, r *http.Request) {
	c := make(chan []Prayer)
	go CrawlBCM(c, 8)
	x := <-c
	fmt.Println(x)

	js, _ := json.Marshal(x)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func handleRequest() {
	http.HandleFunc("/", processNAMAZ)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequest()
}
