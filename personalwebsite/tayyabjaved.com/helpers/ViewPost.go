package helpers

import (
	"context"
	"log"
	entity "microsomes/tinyscrapes/bcm/mod/entitiy"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func ViewPost(postid string, c chan entity.PostR) {
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

	r, _ := client.Collection("posts").Doc(postid).Get(context.Background())

	var p entity.Post
	r.DataTo(&p)

	pp := entity.PostR{
		Title: p.Title,
		Body:  p.Body,
		Date:  p.Date,
		Unix:  p.Unix,
		DocID: r.Ref.ID,
	}

	c <- pp

}
