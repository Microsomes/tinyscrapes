package helpers

import (
	"context"
	"log"
	entity "microsomes/tinyscrapes/bcm/mod/entitiy"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func CreatePost(p *entity.Post, c chan string) {

	if p.Title == "" {
		c <- "0"
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
	ref, _, err := client.Collection("posts").Add(context.Background(), p)
	if err != nil {
		c <- "0"
		return
	}
	c <- ref.ID
}
