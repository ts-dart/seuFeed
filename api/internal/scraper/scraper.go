package scraper

import (
	"fmt"
	"time"
)

type Post struct {
	PostHrefLink string
	PostImgSrc   string
	PostText string
	Font string
	FontImgSrc string
	Section string
}

var PostsList []Post

func Scraper() {
	PostsList = make([]Post, 0, 100)
	PostsList = append(PostsList, urlo()...)
	PostsList = append(PostsList, urlt()...)

	for {
		fmt.Println(PostsList)
		time.Sleep(3 * time.Minute)
	} 
}