package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	entity "microsomes/tinyscrapes/bcm/mod/entitiy"
	helpers "microsomes/tinyscrapes/bcm/mod/helpers"
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

	var d = time.Now().Day()

	var newD = fmt.Sprintf("%02d", d) // Prints to stdout '000012'

	c := make(chan []mosquescrappers.Prayer)
	go mosquescrappers.CrawlBCM(c, month, strconv.Itoa(time.Now().Year()))
	x := <-c

	_ = x

	var res = mosquescrappers.Prayer{}

	for _, val := range x {
		var cuday = strings.Split(val.Month, " ")[1]
		if (cuday) == newD {
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
	tepl, _ := template.ParseFiles("templates/tj/newindex.html")
	c := make(chan []entity.PostR)
	go helpers.GetAllPosts(c)
	x := <-c
	type toRe struct {
		Data []entity.PostR
	}
	xo := toRe{
		Data: x,
	}

	tepl.Execute(w, xo)
}

func handleCreate(w http.ResponseWriter, h *http.Request) {
	tepl, _ := template.ParseFiles("templates/tj/create.html")
	tepl.Execute(w, "")
}

func processCreate(w http.ResponseWriter, h *http.Request) {
	h.ParseForm()
	title := h.PostForm.Get("title")
	body := h.PostForm.Get("body")

	p := entity.Post{
		Title: title,
		Body:  body,
		Date:  time.Now(),
		Unix:  int(time.Now().Unix()),
	}
	c := make(chan string)
	go helpers.CreatePost(&p, c)
	x := <-c
	// http.Redirect(w, h, fmt.Sprintf("/view?postid=%s", x), 200)
	http.Redirect(w, h, fmt.Sprintf("/view?postid=%s", x), http.StatusSeeOther)

	///fmt.Fprintf(w, x)

}

func processPostView(w http.ResponseWriter, h *http.Request) {
	var postid = h.URL.Query().Get("postid")
	if postid == "" {
		fmt.Fprintf(w, "no ?postid present")
		return
	}
	c := make(chan entity.PostR)
	go helpers.ViewPost(postid, c)
	x := <-c

	tepl, _ := template.ParseFiles("templates/tj/view.html")
	tepl.Execute(w, x)

}

func processEditPost(w http.ResponseWriter, h *http.Request) {
	var postid = h.URL.Query().Get("postid")
	if postid == "" {
		fmt.Fprint(w, "No ?postid found")
	}

	m := make(chan entity.PostR)
	go helpers.ViewPost(postid, m)
	x := <-m

	tepl, _ := template.ParseFiles("templates/tj/edit.html")
	tepl.Execute(w, x)

}

func processUpdate(w http.ResponseWriter, h *http.Request) {
	h.ParseForm()
	title := h.PostForm.Get("title")
	body := h.PostForm.Get("body")
	postid := h.PostForm.Get("documentid")
	if postid == "" {
		fmt.Fprint(w, "NO postid present")
		return
	}
	m := make(chan entity.PostR)
	go helpers.ViewPost(postid, m)
	x := <-m
	if x.Title == "" {
		fmt.Fprint(w, "invalid postid, post not found")
		return
	}

	newPost := entity.PostR{
		Title: title,
		Body:  body,
		DocID: postid,
	}

	mm := make(chan entity.PostR)

	go helpers.UpdatePost(newPost, mm)
	res := <-mm
	_ = res
	http.Redirect(w, h, fmt.Sprintf("/edit?postid=%s", newPost.DocID), http.StatusSeeOther)

}

func handleCV(w http.ResponseWriter, h *http.Request) {
	tepl, _ := template.ParseFiles("templates/tj/static/tayyabcv.html")
	tepl.Execute(w, "")
}

func handlePortfolio(w http.ResponseWriter, h *http.Request) {

	type PortfolioContent struct {
		Title   string
		Subline string
	}

	type PortfolioData struct {
		HeaderImage  string
		Name         string
		Technologies []string
		Slug         string
		Extra        PortfolioContent
		SampleLink   string
		IsSample     bool
	}

	allPort := []PortfolioData{}

	allPort = append(allPort, PortfolioData{
		HeaderImage:  "https://sotechnology.co.uk/wp-content/uploads/2021/04/BD_feature@1x.jpg",
		Name:         "Bespoke Cake Editor",
		Technologies: []string{"VUE JS", "NODE JS", "SVG"},
		Slug:         "bespoke-cake-editor",
		SampleLink:   "https://upbeat-curie-5573c6.netlify.app/?test=dd",
		IsSample:     true,
		Extra: PortfolioContent{
			Title:   "Cool little Cake Editor",
			Subline: "A commercial project, for Bakerdays.com 6 month- Development Time",
		},
	})

	allPort = append(allPort, PortfolioData{
		HeaderImage:  "https://sotechnology.co.uk/wp-content/uploads/2021/04/BD_feature@1x.jpg",
		Name:         "Bespoke Cake Editor",
		Technologies: []string{"VUE JS", "NODE JS", "SVG"},
		Slug:         "bespoke-cake-editor",
		SampleLink:   "https://upbeat-curie-5573c6.netlify.app/?test=dd",
		IsSample:     false,
		Extra: PortfolioContent{
			Title:   "Cool little Cake Editor",
			Subline: "A commercial project, for Bakerdays.com 6 month- Development Time",
		},
	})

	type PorfolioPage struct {
		PageTitle string
		Port      []PortfolioData
	}

	tepl, _ := template.ParseFiles("templates/tj/portfolio.html")
	tepl.Execute(w, &PorfolioPage{
		PageTitle: "Portfolio/Work",
		Port:      allPort,
	})
}

func handleContact(w http.ResponseWriter, h *http.Request) {
	tepl, _ := template.ParseFiles("templates/tj/contact.html")
	tepl.Execute(w, "")
}

func handleBlog(w http.ResponseWriter, h *http.Request) {
	tepl, _ := template.ParseFiles("templates/tj/blog.html")
	tepl.Execute(w, "")
}

func openCountries() (*os.File, bool) {
	jsonFile, err := os.Open("countries.json")
	var isError bool
	if err != nil {
		isError = true
	} else {
		isError = false
	}
	return jsonFile, isError
}

func handleCity(w http.ResponseWriter, h *http.Request) {
	country := h.URL.Query().Get("country")
	jsonFile, isError := openCountries()
	if isError {
		fmt.Fprint(w, "Something went wrong try again later")
		return
	}
	fmt.Println("Successfully Opened users.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var d map[string][]string

	err2 := json.Unmarshal(byteValue, &d)
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	type City struct {
		Country string
		Name    string
	}

	var foundCities []City

	for i, val := range d {
		if i == country {
			for _, v := range val {
				foundCities = append(foundCities, City{Name: v, Country: country})
			}
		}
	}

	js, _ := json.Marshal(foundCities)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func handleCountry(w http.ResponseWriter, h *http.Request) {
	jsonFile, isError := openCountries()
	if isError {
		fmt.Fprint(w, "Something went wrong try again later")
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var d map[string]interface{}
	json.Unmarshal(byteValue, &d)

	type Countries struct {
		CountryName string
	}

	var allCountries []Countries

	for country := range d {
		allCountries = append(allCountries, Countries{CountryName: country})
	}

	js, _ := json.Marshal(allCountries)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func handleRequest() {
	//all api calls return json
	http.HandleFunc("/", handleHomePage)

	//for the weather app
	http.HandleFunc("/cities", handleCity)
	http.HandleFunc("/countries", handleCountry)

	http.HandleFunc("/cv", handleCV)
	http.HandleFunc("/portfolio", handlePortfolio)
	http.HandleFunc("/contact", handleContact)
	http.HandleFunc("/create", handleCreate)
	http.HandleFunc("/processCreate", processCreate)
	http.HandleFunc("/processUpdate", processUpdate)
	http.HandleFunc("/blog", handleBlog)
	http.HandleFunc("/view", processPostView)
	http.HandleFunc("/edit", processEditPost)
	http.HandleFunc("bcmmonth/", processNAMAZ)
	http.HandleFunc("/bcmall", showAllYear)
	http.HandleFunc("/bcmc", currentBCM)
	//
	//these generate templates
	http.HandleFunc("/bcm", handleBCM)
	http.HandleFunc("/isna", handleISNA)
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "10000"
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/tj/static"))))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}

func main() {
	handleRequest()
}
