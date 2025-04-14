export interface PostsList {
	postHrefLink: string, 
	postImgSrc: string, 
	postText: string,
	font: string,
	fontImgSrc: string, 
	section: string,
}

export interface CurrentClimateData {
	time: string,
	interval: number,
	temperature2m: number,
	windSpeed10m: number,
}