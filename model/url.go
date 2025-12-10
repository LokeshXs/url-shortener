package model

import (
	"errors"
	"time"

	"github.com/Lokeshxs/url-shortener/db"
	"github.com/Lokeshxs/url-shortener/utils"

)

type URL struct {
	LongURL   string `json:"long_url"`
	ShortURL  string
	ShortCode string     `json:"shortcode"`
	ExpiredAT *time.Time `json:"expired_at"`
}
type ResultURL struct {
	OriginalURL string     `json:"original_url"`
	ShortURL    string     `json:"shorturl"`
	Clicks      int64      `json:"clicks"`
	ExpiredAT   *time.Time `json:"expired_at"`
}

func (url *URL) ShortenURL(userId string, expired_at *time.Time, customShortCode string,baseURL string) error {

	var shortCode string

	// Check for custom short code in DB if its already taken on assigned
	if len(customShortCode) > 0 {
		if len(customShortCode) < 6 {

			return errors.New("Short code length should be minimum 6")
		}
		query := `SELECT 1 FROM urls WHERE shortcode = $1`

		resultedRow, err := db.DB.Query(query, customShortCode)

		if err != nil {
			
			return err
		}

		resultedRow.Close()

		if resultedRow.Next() {
			return errors.New("Shortcode is already taken")
		}

		shortCode = customShortCode
	}

	// PATTERN :- /mydomain/:shortcode (shortcode: 6chars long with alphanumeric characters and case sensitivity)

	if len(shortCode) == 0 {

		shortCode = utils.GenerateShortCode(utils.SHORT_CODE_LENGTH)
	}

	shortUrl := baseURL + "/" + shortCode

	// Saving the data in DB
	query := `
	INSERT INTO urls (original_url, shortcode, user_id, expired_at)
	VALUES ($1, $2, $3, $4);
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(url.LongURL, shortCode, userId, expired_at)

	if err != nil {
		return err
	}

	// Assigning the short url to the struct

	url.ShortURL = shortUrl

	return nil

}


func DeleteURL(shortCode string)(error){

	query := `DELETE FROM urls WHERE shortcode = $1;`;

	_,err := db.DB.Exec(query,shortCode);

	if(err != nil){

		return err;
	}


	return  nil;


}

func GetURL(shortCode string) (string, error) {

	query := `
	SELECT original_url, expired_at FROM urls 
	WHERE shortcode = $1;

	`

	result := db.DB.QueryRow(query, shortCode)

	var original_url string
	var expired_at *time.Time
	err := result.Scan(&original_url, &expired_at)
	if err != nil {

		return "", err
	}
	

	currentTime := time.Now()

	if expired_at != nil {

		if currentTime.Unix() > (*expired_at).Unix() {
			return "", errors.New("Short URL is expired!")
		}
	}

	// update the clicks on url
	query = `UPDATE urls SET clicks = clicks + 1
	WHERE shortcode = $1;
	`

	// Not returning the error because this is a not must query to run.
	db.DB.Exec(query, shortCode)

	return original_url, nil

}

func GetAllURLS(userId string, pageNumber int64, limit int64,baseURL string) ([]ResultURL, int64, error) {

	query := `
	SELECT original_url, shortcode, clicks,  expired_at 
	FROM urls
	WHERE user_id = $1
	ORDER BY id LIMIT $2 OFFSET $3;
	`

	offset := (pageNumber - 1) * limit

	rows, err := db.DB.Query(query, userId, limit, offset)

	if err != nil {

		return nil, 0, err
	}

	defer rows.Close()

	var result = []ResultURL{}

	for rows.Next() {
		var url ResultURL

		err = rows.Scan(&url.OriginalURL, &url.ShortURL, &url.Clicks, &url.ExpiredAT)

		url.ShortURL = baseURL + "/" + url.ShortURL

		if err != nil {

			return nil, 0, err
		}

		result = append(result, url)
	}

	// Get the total count

	var total int64
	query = `SELECT COUNT(*) FROM urls WHERE user_id = $1;`

	err = db.DB.QueryRow(query,userId).Scan(&total)

	if err != nil {

		return nil, 0, err
	}

	// fmt.Println(result);

	return result, total, nil
}
