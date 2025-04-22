export interface Post {
	post_href_link: string, 
	post_img_src: string, 
	post_text: string,
	font: string,
	font_img_src: string, 
	section: string,
}

export interface CurrentClimateData {
	time: string,
	interval: number,
	temperature2m: number,
	windSpeed10m: number,
}