package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"unicode/utf8"

	firebase "firebase.google.com/go"
	"github.com/gocolly/colly"
	"google.golang.org/api/option"
)

//function is responsible for saving the cache
func SaveToCache(Prayers []Prayer, monthSaved int, status chan int) {
	fmt.Println("saving to cache")
	opt := option.WithCredentialsFile("firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("error initializing app:")
		return
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	client.Collection("cache").Add(context.Background(), map[string]interface{}{
		"status": "ok",
		"month":  monthSaved,
		"all":    Prayers,
	})

	time.Sleep(10 * time.Second)
	status <- 3
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
			var fajrJamat = ""
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
				if i == 3 {
					fajrJamat = h.Text
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
					Month:     month,
					Day:       day,
					Fajr:      fajr,
					FajrJamat: fajrJamat,
					Sunrise:   sunrise,
					Zuhr:      zuhr,
					Asr:       asr,
					Maghrib:   maghrib,
					Isha:      isha,
				})
			}
		})

		//save to cache first
		cindex := make(chan int)
		go SaveToCache(Prayers, monthRequested, cindex)

		ltindex := <-cindex
		fmt.Println(ltindex)

		cc <- Prayers

	})
	c.Visit(url)
}
