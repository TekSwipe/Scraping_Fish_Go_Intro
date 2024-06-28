package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	params := url.Values{}
	params.Add("api_key", "[your API key]")
	params.Add("url", "https://example.com")

	resp, err := http.Get(fmt.Sprintf("%s?%s", "https://scraping.narf.ai/api/v1/", params.Encode()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}
