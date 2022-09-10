package utils

import (
	"fmt"
	"net/http"
	"log"
	"bytes"
)

func HtbGet(url string) *http.Response {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	jwt_token := GetConfigValue("token")
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", "HTB-Tool")
	req.Header.Set("Authorization", "Bearer " + jwt_token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

func HtbPost(url string, jsonData []byte) *http.Response {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	jwt_token := GetConfigValue("token")
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", "HTB-Tool")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + jwt_token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}