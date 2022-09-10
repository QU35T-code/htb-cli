package utils

import (
	"io"
	"encoding/json"
	"log"
	"net/http"
)

func SearchMachineIDByName(machine_name string) string {
	url := "https://www.hackthebox.com/api/v4/search/fetch?query=" + machine_name
	resp := HtbGet(url)
	json_body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(json_body), &result)
	machinesArray := result["machines"].([]interface{})
	machineData := machinesArray[0].(map[string]interface{})
	return machineData["id"].(string)
}

func ParseJsonMessage(resp *http.Response, key string) interface{} {
	json_body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(json_body), &result)
    return result[key]
}