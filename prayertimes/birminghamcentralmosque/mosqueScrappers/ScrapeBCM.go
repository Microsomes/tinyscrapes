package mosquescrappers

import (
	"context"
	"fmt"
	"log"
	"unicode/utf8"

	firebase "firebase.google.com/go"
	"github.com/gocolly/colly"
	"google.golang.org/api/option"
)

type toSave struct {
	Status  string
	Prayers []Prayer
}

//function is responsible for saving the cache
func SaveToCache(Prayers []Prayer, monthSaved int, status chan int, cacheKey string) {
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

	var toL = toSave{
		Status:  "OK",
		Prayers: Prayers,
	}

	_, err = client.Collection("cache").Doc(fmt.Sprintf("cache_%s_%d", cacheKey, monthSaved)).Set(
		context.Background(),
		toL,
	)
	if err != nil {
		status <- 0
	}

	status <- 3
}

func findCache(cacheKey string, monthRequested int, c chan toSave) {
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

	dsnap, err := client.Collection("cache").Doc(fmt.Sprintf("cache_%s_%d", cacheKey, monthRequested)).Get(context.Background())
	if err != nil {
		c <- toSave{Status: "err"}
	}
	if dsnap.Exists() {

		var dd toSave

		dsnap.DataTo(&dd)
		c <- dd

	} else {
		c <- toSave{Status: "err"}
	}

}

type Prayer struct {
	Day       string `firestore:"Day,omitempty"`
	Month     string `firestore:"Month,omitempty"`
	Fajr      string `firestore:"Fajr,omitempty"`
	FajrJamat string `firestore:"FajrJamat,omitempty"`
	Sunrise   string `firestore:"Sunrise,omitempty"`
	Zuhr      string `firestore:"Zuhr,omitempty"`
	Asr       string `firestore:"Asr,omitempty"`
	Isha      string `firestore:"Isha,omitempty"`
	Maghrib   string `firestore:"Maghrib,omitempty"`
}

func CrawlBCM(cc chan []Prayer, monthRequested int, cacheKey string) {

	findCacheChannel := make(chan toSave)
	go findCache(cacheKey, monthRequested, findCacheChannel)
	didFindCache := <-findCacheChannel

	if didFindCache.Status != "err" {
		fmt.Println("cache found")
		cc <- didFindCache.Prayers
		return
	}

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
		go SaveToCache(Prayers, monthRequested, cindex, cacheKey)

		ltindex := <-cindex
		fmt.Println(fmt.Sprintf("Cache save result:%d", ltindex))

		cc <- Prayers

	})
	c.Visit(url)
}
