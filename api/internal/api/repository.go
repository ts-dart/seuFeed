package api

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"io"
)

type Post struct {
	PostHrefLink string `json:"post_href_link"`
	PostImgSrc string `json:"post_img_src"`
	PostText string `json:"post_text"`
	Font string `json:"font"`
	FontImgSrc string `json:"font_img_src"`
	Section string `json:"section"`
}

func respository() ([]Post) {
	var posts []Post
	readJson(&posts)

	fmt.Println(posts)
	return posts
}

func readJson(posts *[]Post) {
	file, err := os.Open("../scraper/data.json")
	if err != nil {
		log.Fatal("Erro ao abrir o arquivo", err)
	}
	defer file.Close()
	
	jsonData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Erro ao ler o arquivo", err)
	}

	err = json.Unmarshal(jsonData, &posts)
	if err != nil {
		log.Fatal("Erro ao ler o arquivo", err)
	}
}