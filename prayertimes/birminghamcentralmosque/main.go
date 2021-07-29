package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	mosquescrappers "microsomes/tinyscrapes/bcm/mod/mosqueScrappers"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func processNAMAZ(w http.ResponseWriter, r *http.Request) {
	var cacheKey = r.URL.Query().Get("cachekey")

	var monthRequested = r.URL.Query().Get("month")

	var mo, err = strconv.Atoi(monthRequested)

	if err != nil {
		fmt.Fprintf(w, "month is wrong ?month=2 ?month only accepts a int")
		return
	}

	if cacheKey != "" {
		//all data is now live
	} else {
		fmt.Fprintf(w, "?cachekey is missing")
		return
	}

	c := make(chan []mosquescrappers.Prayer)
	go mosquescrappers.CrawlBCM(c, mo, cacheKey)
	x := <-c
	fmt.Println(x)

	js, _ := json.Marshal(x)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func showAllYear(w http.ResponseWriter, h *http.Request) {

	var cacheKey = h.URL.Query().Get("cachekey")

	if cacheKey == "" {
		fmt.Fprint(w, "No ?cachekey present")
		return
	}

	cch := make(chan [][]mosquescrappers.Prayer)
	go mosquescrappers.BCMYear(cacheKey, cch)
	res := <-cch

	w.Header().Set("Content-Type", "application/json")
	js, _ := json.Marshal(res)

	w.Write(js)

}

type StrucutedResponse struct {
	Status string
	Msg    string
	Data   interface{}
}

func handleISNA(w http.ResponseWriter, h *http.Request) {
	m := make(chan []mosquescrappers.ISNAPrayer)
	go mosquescrappers.ScrapeISNACanada(m)
	res := <-m
	fmt.Println(res)

	var toReturn = StrucutedResponse{
		Status: "OK",
		Msg:    "Current Namaz times from https://isnacanada.com/",
		Data:   res,
	}

	tepl, _ := template.ParseFiles("templates/isna/index.html")

	tepl.Execute(w, toReturn)

	// w.Write(jss)
}

func currentBCMLogic(resc chan mosquescrappers.Prayer) {
	var month = int(time.Now().Month())

	c := make(chan []mosquescrappers.Prayer)
	go mosquescrappers.CrawlBCM(c, month, strconv.Itoa(time.Now().Year()))
	x := <-c

	var res = mosquescrappers.Prayer{}

	for _, val := range x {
		var cuday = strings.Split(val.Month, " ")[1]
		if (cuday) == strconv.Itoa((time.Now().Day())) {
			res = val
		}
	}
	resc <- res
}

func currentBCM(w http.ResponseWriter, h *http.Request) {

	c := make(chan mosquescrappers.Prayer)
	go currentBCMLogic(c)
	res := <-c

	js, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func handleBCM(w http.ResponseWriter, h *http.Request) {
	c := make(chan mosquescrappers.Prayer)
	go currentBCMLogic(c)
	res := <-c

	type ISNAPrayer struct {
		PrayerName  string
		PrayerBegin string
		PrayerJamah string
		HijriDate   string
		Gregorian   string
	}

	var cd []ISNAPrayer

	cd = append(cd, ISNAPrayer{
		PrayerName:  "Fajr",
		PrayerBegin: res.Fajr,
		PrayerJamah: res.FajrJamat,
		Gregorian:   res.Month,
	})
	cd = append(cd, ISNAPrayer{
		PrayerName:  "Zuhr",
		PrayerBegin: res.Zuhr,
		PrayerJamah: res.ZuhrJamat,
	})
	cd = append(cd, ISNAPrayer{
		PrayerName:  "Asr",
		PrayerBegin: res.Asr,
		PrayerJamah: res.AsrJamat,
	})
	cd = append(cd, ISNAPrayer{
		PrayerName:  "Maghrib",
		PrayerBegin: res.Maghrib,
		PrayerJamah: res.MaghribJamat,
	})
	cd = append(cd, ISNAPrayer{
		PrayerName:  "Isha",
		PrayerBegin: res.Isha,
		PrayerJamah: res.IshaJamat,
	})

	var toReturn = StrucutedResponse{
		Status: "OK",
		Msg:    "Current Namaz times from https://centralmosque.org.uk/",
		Data:   cd,
	}

	tepl, _ := template.ParseFiles("templates/bcm/index.html")
	tepl.Execute(w, toReturn)

}

func handleHomePage(w http.ResponseWriter, h *http.Request) {
	tepl, _ := template.ParseFiles("templates/tj/index.html")
	tepl.Execute(w, "")
}

func handleCreate(w http.ResponseWriter, h *http.Request) {
	tepl, _ := template.ParseFiles("templates/tj/create.html")
	tepl.Execute(w, "")
}

func handleRequest() {
	//all api calls return json
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/create", handleCreate)
	http.HandleFunc("bcmmonth/", processNAMAZ)
	http.HandleFunc("/bcmall", showAllYear)
	http.HandleFunc("/bcmc", currentBCM)

	//these generate templates
	http.HandleFunc("/bcm", handleBCM)
	http.HandleFunc("/isna", handleISNA)
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "10000"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}

func main() {
	year, month, day := time.Now().Date()
	fmt.Println(year)
	fmt.Println(int(month))
	fmt.Println(day)

	handleRequest()

}
