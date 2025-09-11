package model

import (
	"github.com/Lokeshxs/url-shortener/db"
	"github.com/Lokeshxs/url-shortener/utils"
)

type URL struct {
	LongURL  string `json:"long_url"`
	ShortURL string
}

func (url *URL) ShortenURL(domain string) error {

	// PATTERN :- /mydomain/:shortcode (shortcode: 6chars long with alphanumeric characters and case sensitivity)

	shortCode := utils.GenerateShortCode(utils.SHORT_CODE_LENGTH)

	shortUrl := "https://" + domain + "/u/" + shortCode

	// Saving the data in DB
	query := `
	INSERT INTO urls (original_url, shortcode, user_id)
	VALUES ($1, $2, $3);
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	//todo: Get the user id from jwt token or something
	user_id := 1

	_, err = stmt.Exec(url.LongURL, shortCode,user_id)

	if err != nil {
		return err
	}

	// Assigning the short url to the struct

	url.ShortURL = shortUrl

	return nil

}
