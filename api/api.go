package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	baseUrl := "http://openapi.molit.go.kr/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcAptTradeDev"
	key := os.Getenv("OPEN_API_KEY")
	params := url.Values{}
	params.Set("serviceKey", key)
	params.Set("pageNo", "1")
	params.Set("numOfRows", "10")
	params.Set("LAWD_CD", "11110")
	params.Set("DEAL_YMD", "201512")

	fullUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())
	resp, err := http.Get(fullUrl)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println("Response:", string(body))
}
