package helpers

import (
	"context"
	"log"
	entity "microsomes/tinyscrapes/bcm/mod/entitiy"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func UpdatePost(newPost entity.PostR, c chan entity.PostR) {
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

	_, err = client.Collection("posts").Doc(newPost.DocID).Set(context.Background(), newPost)
	c <- newPost
}
