package api

func postsBySection(ft string) ([]Post) {
	r := respository()
	var filteredPosts []Post

	for _, p := range r {
		if p.section == ft {
			filteredPosts = append(filteredPosts, p)
		}
	}

	return filteredPosts
}
