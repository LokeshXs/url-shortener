package model

import "github.com/Lokeshxs/url-shortener/utils"

type URL struct{
	LongURL string `json:"long_url"`
	ShortURL string
}


func (url *URL) ShortenURL(domain string){

	// PATTERN :- /mydomain/:shortcode (shortcode: 6chars long with alphanumeric and case sensitivity)

	shortCode := utils.GenerateShortCode(utils.SHORT_CODE_LENGTH);

	shortUrl := "https://"+domain+"/u/"+shortCode;


	url.ShortURL = shortUrl;


}