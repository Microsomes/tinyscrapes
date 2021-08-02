package helpers

import (
	"context"
	"log"
	entity "microsomes/tinyscrapes/bcm/mod/entitiy"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func GetAllPosts(p chan []entity.PostR) {

	var data []entity.PostR

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
	iter := client.Collection("posts").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return
		}
		var ee entity.Post
		doc.DataTo(&ee)

		var newEE = entity.PostR{
			Title: ee.Title,
			Body:  ee.Body,
			Date:  ee.Date,
			DocID: doc.Ref.ID,
		}

		data = append(data, newEE)
	}
	p <- data

}
