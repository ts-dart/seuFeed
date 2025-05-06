package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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

func getClimateDate(latStr string, lonStr string) CurrentClimateData {
	err := godotenv.Load()
	if err != nil {
		panic("Erro ao carregar o arquivo .env")
	}

	baseURL := os.Getenv("OPEN_METEO_URL")
	if baseURL == "" {
		panic("OPEN_METEO_URL não definida no .env")
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		panic(fmt.Sprintf("Latitude inválida: %v", err))
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		panic(fmt.Sprintf("Longitude inválida: %v", err))
	}

	url := fmt.Sprintf("%s&latitude=%.2f&longitude=%.2f&current=temperature_2m,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m", baseURL, lat, lon)

	resp, err := http.Get(url)
	if err != nil {
		panic("Erro ao requisitar os dados da API open-meteo")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Erro ao ler os dados da API open-meteo")
	}

	var apiResp ClimateData
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		panic("Erro ao decodificar o JSON para a estrutura")
	}

	return apiResp.Current
}