package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v3"
)

type APIerror struct {
	Body string `json:"errorMessage"`
	Code int    `json:"errorCode"`
}

func Hello(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello World!"})
}

func GetName(c fiber.Ctx) error {
	baseURL := os.Getenv("NCP_API_URL")
	koreanName := c.Queries()["name"]
	rawURL := baseURL + url.QueryEscape(koreanName)
	reqURL, err := url.Parse(rawURL)
	if err != nil {
		log.Fatal(err)
	}
	req := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"X-NCP-APIGW-API-KEY-ID": {os.Getenv("NCP_API_KEY_ID")},
			"X-NCP-APIGW-API-KEY":    {os.Getenv("NCP_API_KEY")}},
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode == http.StatusOK {
		return c.Status(res.StatusCode).JSON(fiber.Map{"body": body})
	} else {
		errRes := APIerror{}
		json.Unmarshal(body, &errRes)
		return c.Status(errRes.Code).JSON(fiber.Map{"body": errRes.Body})
	}
}
