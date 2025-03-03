package api

func postsBySection(s string) ([]Post) {
	r := respository()
	var filteredPosts []Post

	for _, p := range r {
		if p.section == s {
			filteredPosts = append(filteredPosts, p)
		}
	}

	return filteredPosts
}
