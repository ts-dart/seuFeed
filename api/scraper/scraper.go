package scraper

import "fmt"

type Post struct {
	PostHrefLink string
	PostImgSrc   string
	PostCategory string
	PostTitle    string
}

func Scraper() {
	postsList := make([]Post, 60, 60)
	Urlo(postsList)
	fmt.Println(postsList)

	/*
	index := 0

	for {
		index += 1
		Urlo(postsList)
	} */
}