package helpers

import (
	"context"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Post struct {
	Title string    `firestore:"Title,omitempty"`
	Body  string    `firestore:"Body,omitempty"`
	Date  time.Time `firestore:"Date,omitempty"`
	Unix  int       `firestore:"Unix,omitempty"`
}

func CreatePost(p *Post, c chan int) {
	if p.Title == "" {
		c <- 0
		return
	}
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
	ref, wri, err := client.Collection("posts").Add(context.Background(), p)
	if err != nil {
		c <- 0
	}
	c <- 1
	time.Sleep(2 * time.Second)
	c <- 2
}
