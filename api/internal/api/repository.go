package api

import (
	"github.com/ts-dart/seuFeed/api/internal/scraper"
	"fmt"
)

type Post struct {
	postHrefLink string
	postImgSrc   string
	postText string
	font string
	fontImgSrc string
	section string
}

func respository() ([]Post) {
	posts := scraper.PostsList
	fmt.Println(posts)
	var convertedPosts []Post

	for _, p := range posts {
		convertedPosts = append(convertedPosts, Post{
			postHrefLink: p.PostHrefLink,
			postImgSrc:   p.PostImgSrc,
			postText:     p.PostText,
			font:         p.Font,
			fontImgSrc:   p.FontImgSrc,
			section:      p.Section,
		})
	}

	return convertedPosts
}