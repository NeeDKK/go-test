package main

import (
	"github.com/gocolly/colly"
	"log"
)

func main() {
	c := colly.NewCollector()

	// authenticate
	err := c.Post("https://10.25.10.120/api/login", map[string]string{"username": "wzj", "password": "Admin@123456", "captcha": "899104", "captchaId": "naTbJVao9fEuM8tHuq3N"})
	if err != nil {
		log.Fatal(err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	// start scraping
	c.Visit("https://10.25.10.120/")
}
