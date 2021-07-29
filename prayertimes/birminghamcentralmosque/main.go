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

func handleRequest() {
	http.HandleFunc("/", processNAMAZ)
	http.HandleFunc("/all", showAllYear)
	http.HandleFunc("/isna", handleISNA)
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "10000"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}

func main() {

	handleRequest()

}
