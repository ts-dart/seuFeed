package scraper

import (
	"fmt"
)

func urlt() {
	url, font, fontImg := envValidator("URLT_ENV", "URLT_ENV_FONT", "URLT_ENV_FONT_IMG")

	fmt.Println(url, font, fontImg)
}
