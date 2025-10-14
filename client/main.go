package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func getText(c *http.Client, url string) (string, error) {
	resp, err := c.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	return string(b), err
}

func postJSON(c *http.Client, url, jsonBody string) (string, error) {
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	return string(b), err
}

func main() {
	base := "http://localhost:8080"
	httpClient := &http.Client{}

	ver, _ := getText(httpClient, base+"/version")
	fmt.Println(strings.TrimSpace(ver))

	body := `{"inputString":"eXl4"}`
	out, _ := postJSON(httpClient, base+"/decode", body)
	var resp struct {
		OutputString string `json:"outputString"`
	}
	_ = json.Unmarshal([]byte(out), &resp)
	fmt.Println(resp.OutputString)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, base+"/hard-op", nil)

	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("false")
		return
	}
	defer res.Body.Close()
	io.Copy(io.Discard, res.Body)
	fmt.Printf("true, %d\n", res.StatusCode)
}
