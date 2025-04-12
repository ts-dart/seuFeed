package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func getPostsBySection(ft string) ([]Post) {
	r := respository()
	var filteredPosts []Post

	for _, p := range r {
		if p.Section == ft {
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

func getClimateDate(latStr string, lonStr string) {
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		panic(fmt.Sprintf("Latitude inválida: %v", err))
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		panic(fmt.Sprintf("Longitude inválida: %v", err))
	}

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.2f&longitude=%.2f&current=temperature_2m,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m", lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("Erro ao requisitar os dados da api open-meteo"))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Erro ao ler os dados da api open-meteo"))
	}

	var apiResp ClimateData
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		panic(fmt.Sprintf("Erro ao decodificar o json para estutura"))
	}

	r := ClimateData{
		Current{
			Time: apiResp.Current.Time, 
			Temperature2m: apiResp.Current.Temperature2m, 
			WindSpeed10m: apiResp.Current.WindSpeed10m, 
			Interval: apiResp.Current.Interval,
		},
	}

	//formatedResp := ClimateData{apiResp.Current.Time, apiResp.Current.Temperature2m, apiResp.Current.WindSpeed10m, apiResp.Current.Interval}
}