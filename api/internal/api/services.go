package api

import (
	"fmt"
)

func getPostsBySection(ft string) ([]Post) {
	r := respository()
	var filteredPosts []Post

	for _, p := range r {
		if p.section == ft {
			filteredPosts = append(filteredPosts, p)
		}
	}

	return filteredPosts
}

func getAllPosts() ([]Post) {
	r := respository()
	fmt.Println()
	return r
}