package entity

import "time"

type Post struct {
	Title string    `firestore:"Title,omitempty"`
	Body  string    `firestore:"Body,omitempty"`
	Date  time.Time `firestore:"Date,omitempty"`
	Unix  int       `firestore:"Unix,omitempty"`
}
