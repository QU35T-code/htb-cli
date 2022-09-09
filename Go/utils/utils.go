package utils

import (
	"fmt"
	"io"
	"encoding/json"
	"log"
)

func SearchMachineIDByName(machine_name string) int {
	fmt.Println(machine_name)
	url := "https://www.hackthebox.com/api/v4/search/fetch?query=" + machine_name
	resp := HtbGet(url)
	json_body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(json_body), &result)
	machines := result["machines"]
	fmt.Println(machines)
	// var result2 map[string]interface{}
	// json.Unmarshal(map[string]interface{}(machines), &result2)
    // fmt.Println(result2)
	return 1
}