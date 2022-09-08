package utils

import (
	"fmt"
	"net/http"
	"log"
)

func HtbGet(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	jwt_token := GetHtbToken()
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

func htb_post() {
}