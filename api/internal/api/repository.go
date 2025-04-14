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
	//Hora
}

type CurrentClimateData struct {
	Time string `json:"time"`
	Interval int `json:"interval"`
	Temperature2m float64 `json:"temperature_2m"`
	WindSpeed10m float64 `json:"wind_speed_10m"`
}

type ClimateData struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	GenerationTimeMS float64 `json:"generationtime_ms"`
	UTCOffsetSeconds int `json:"utc_offset_seconds"`
	Timezone string  `json:"timezone"`
	TimezoneAbbreviation string `json:"timezone_abbreviation"`
	Elevation float64 `json:"elevation"`

	CurrentUnits struct {
		Time string `json:"time"`
		Interval string `json:"interval"`
		Temperature2m string `json:"temperature_2m"`
		WindSpeed10m string `json:"wind_speed_10m"`
	} `json:"current_units"`

	Current struct {
		Time string `json:"time"`
		Interval int `json:"interval"`
		Temperature2m float64 `json:"temperature_2m"`
		WindSpeed10m float64 `json:"wind_speed_10m"`
	} `json:"current"`

	HourlyUnits struct {
		Time string `json:"time"`
		Temperature2m string `json:"temperature_2m"`
		RelativeHumidity2m string `json:"relative_humidity_2m"`
		WindSpeed10m string `json:"wind_speed_10m"`
	} `json:"hourly_units"`

	Hourly struct {
		Time []string `json:"time"`
		Temperature2m []float64 `json:"temperature_2m"`
		RelativeHumidity2m []int `json:"relative_humidity_2m"`
		WindSpeed10m []float64 `json:"wind_speed_10m"`
	} `json:"hourly"`
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