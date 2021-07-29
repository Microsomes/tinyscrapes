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
	"sync"
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

func worker(id int, wg *sync.WaitGroup, toReturn chan []mosquescrappers.Prayer) {
	c := make(chan []mosquescrappers.Prayer)
	go mosquescrappers.CrawlBCM(c, id, "2021")
	x := <-c
	toReturn <- x
	wg.Done()

}

func showAllYear(w http.ResponseWriter, h *http.Request) {
	var wg sync.WaitGroup

	var cacheKey = h.URL.Query().Get("cachekey")

	if cacheKey == "" {
		fmt.Fprint(w, "No ?cachekey present")
		return
	}

	c := make(chan []mosquescrappers.Prayer)

	var allResults [][]mosquescrappers.Prayer

	for i := 1; i <= 12; i++ {
		wg.Add(1)
		go worker(i, &wg, c)
		res := <-c
		allResults = append(allResults, res)
	}
	wg.Wait()

	w.Header().Set("Content-Type", "application/json")
	js, _ := json.Marshal(allResults)

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

	// client.Collection("test").Add(context.Background(), map[string]interface{}{
	// 	"msg": "Hello",
	// })

	handleRequest()

	// fmt.Println("saving to cache")
	// opt := option.WithCredentialsFile("firebase.json")
	// app, err := firebase.NewApp(context.Background(), nil, opt)
	// if err != nil {
	// 	log.Fatal("error initializing app:")
	// 	return
	// }

	// client, err := app.Firestore(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Close()

	// // prayer := mosquescrappers.Prayer{
	// // 	Fajr:      "12",
	// // 	FajrJamat: "12",
	// // 	Sunrise:   "12",
	// // 	Zuhr:      "12",
	// // 	Asr:       "12",
	// // 	Maghrib:   "12",
	// // 	Isha:      "12",
	// // 	Day:       "12",
	// // 	Month:     "12",
	// // }

	// // prayers := []mosquescrappers.Prayer{
	// // 	prayer,
	// // 	prayer,
	// // }

	// // daa := toSave{
	// // 	Status:  "OK",
	// // 	Prayers: prayers,
	// // }

	// // client.Collection("cache").Doc("tocache02").Set(context.Background(), daa)

	// // dsnap, _ := client.Collection("cache").Doc("tocache02").Get(context.Background())

	// // var dataCustom toSave

	// // dsnap.DataTo(&dataCustom)

	// // fmt.Println(dataCustom.Prayers[0].Day)

}
