package scraper

import (
	"fmt"
	"time"
)

type Post struct {
	postHrefLink string
	postImgSrc   string
	postCategory string
	font string
	fontImgSrc string
	section string
}

func Scraper() {
	postsList := make([]Post, 0, 100)
	postsList = append(postsList, urlo()...)
	urlt()

	for {
		fmt.Println(postsList)
		time.Sleep(3 * time.Minute)
	} 
}