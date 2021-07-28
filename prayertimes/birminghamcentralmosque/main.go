package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"unicode/utf8"

	"github.com/gocolly/colly"
)

type Prayer struct {
	Month     string
	Day       string
	Sunrise   string
	Fajr      string
	FajrJamat string
	Zuhr      string
	Asr       string
	Maghrib   string
	Isha      string
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
		cc <- Prayers

	})
	c.Visit(url)
}

func processNAMAZ(w http.ResponseWriter, r *http.Request) {
	var cacheKey = r.URL.Query().Get("cachekey")

	var monthRequested = r.URL.Query().Get("month")

	var mo, err = strconv.Atoi(monthRequested)

	if err != nil {
		fmt.Fprintf(w, "month is wrong ?month=2 ?month only accepts a int")
		return
	}

	if cacheKey != "" {
		dat, err := ioutil.ReadFile(fmt.Sprintf("cache/data_%s_%d.json", cacheKey, mo))
		if err != nil {
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(dat)
			return
		}
	}

	c := make(chan []Prayer)
	go CrawlBCM(c, mo)
	x := <-c
	fmt.Println(x)

	js, _ := json.Marshal(x)
	_ = ioutil.WriteFile(fmt.Sprintf("cache/data_%s_%d.json", cacheKey, mo), js, 0644)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func worker(id int, wg *sync.WaitGroup, toReturn chan []Prayer) {

	dat, err := ioutil.ReadFile(fmt.Sprintf("cache/data_%d_%d.json", 2021, id))

	data := []Prayer{}

	_ = json.Unmarshal([]byte(dat), &data)

	if err != nil {

		c := make(chan []Prayer)

		go CrawlBCM(c, id)
		x := <-c
		js, _ := json.Marshal(x)
		_ = ioutil.WriteFile(fmt.Sprintf("cache/data_%d_%d.json", 2021, id), js, 0644)
		d2 := []Prayer{}
		dat, _ := ioutil.ReadFile(fmt.Sprintf("cache/data_%d_%d.json", 2021, id))

		_ = json.Unmarshal([]byte(dat), &d2)
		toReturn <- d2
		wg.Done()

	} else {
		toReturn <- data
		wg.Done()
		return
	}
}

func showAllYear(w http.ResponseWriter, h *http.Request) {
	var wg sync.WaitGroup

	var cacheKey = h.URL.Query().Get("cachekey")

	if cacheKey == "" {
		fmt.Fprint(w, "No ?cachekey present")
		return
	}

	dat, err := ioutil.ReadFile(fmt.Sprintf("cache/cache_%s.json", cacheKey))

	if err == nil {
		var allResults [][]Prayer

		_ = json.Unmarshal([]byte(dat), &allResults)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dat)
		return
	}

	c := make(chan []Prayer)

	var allResults [][]Prayer

	for i := 1; i <= 12; i++ {
		wg.Add(1)
		go worker(i, &wg, c)
		res := <-c
		allResults = append(allResults, res)
	}
	wg.Wait()

	fmt.Println(allResults)

	w.Header().Set("Content-Type", "application/json")
	js, _ := json.Marshal(allResults)

	ioutil.WriteFile(fmt.Sprintf("cache/cache_%s.json", cacheKey), js, 0644)
	w.Write(js)

}

func handleRequest() {
	http.HandleFunc("/", processNAMAZ)
	http.HandleFunc("/all", showAllYear)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequest()
}
