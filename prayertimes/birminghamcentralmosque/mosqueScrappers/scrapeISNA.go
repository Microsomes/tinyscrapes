package mosquescrappers

import (
	"fmt"

	"github.com/gocolly/colly"
)

type ISNAPrayer struct {
	PrayerName  string
	PrayerBegin string
	PrayerJamah string
	HijriDate   string
	Gregorian   string
}

func ScrapeISNACanada(prayersC chan []ISNAPrayer) {
	c := colly.NewCollector()
	var prayers []ISNAPrayer
	// Find and visit all links
	c.OnHTML("table", func(e *colly.HTMLElement) {
		var hijriDate = e.DOM.Find(".hijriDate").Text()
		var geo = ""

		e.ForEach("tr", func(lt int, h *colly.HTMLElement) {
			if lt == 0 {
				h.ForEach("th", func(i int, h *colly.HTMLElement) {
					if i == 0 {
						geo = h.Text
					}
				})
			}
			var prayerName = h.DOM.Find(".prayerName").Text()
			var prayerBegins = h.DOM.Find(".begins").Text()
			var prayerJamah = h.DOM.Find(".jamah").Text()

			if prayerBegins != "" {
				prayers = append(prayers, ISNAPrayer{
					PrayerName:  prayerName,
					PrayerBegin: prayerBegins,
					PrayerJamah: prayerJamah,
					HijriDate:   hijriDate,
					Gregorian:   geo,
				})
			}
		})
		prayersC <- prayers

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://isnacanada.com/")
}
