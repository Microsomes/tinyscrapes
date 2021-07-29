package helpers

import "fmt"

type Post struct {
	String title
	String body
}

func CreatePost(p *Post) {
	fmt.Println("hello")
}
