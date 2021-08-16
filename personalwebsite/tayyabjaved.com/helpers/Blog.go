package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type BlogResult struct {
	Data struct {
		User struct {
			Name        string `json:"name"`
			BlogHandle  string `json:"blogHandle"`
			Publication struct {
				Posts []struct {
					Title      string    `json:"title"`
					Slug       string    `json:"slug"`
					DateAdded  time.Time `json:"dateAdded"`
					DatePretty string
				} `json:"posts"`
			} `json:"publication"`
		} `json:"user"`
	} `json:"data"`
}

/*
This will get all blogs
TODO add some pagination logic because currently their is none
//although hashnode may add its own pagination
*/
func GetBlogs() BlogResult {
	url := "https://api.hashnode.com/"

	payload := strings.NewReader("{\"query\":\"# Write your query or mutation here\\n{\\n  \\n  user(username:\\\"microsomes\\\"){\\n    name\\n    blogHandle\\n    publication{\\n      posts(page:0){\\n        title\\n        slug\\n        dateAdded\\n      }\\n    }\\n  }\\n  \\n}\\n\"}")
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	//TODO add error handing
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var objMap BlogResult

	err := json.Unmarshal(body, &objMap)
	if err != nil {
		log.Fatal()
	}

	var objMapRef = &objMap

	for _, v := range objMapRef.Data.User.Publication.Posts {
		v.DatePretty = "egoga"
	}

	return objMap
}
