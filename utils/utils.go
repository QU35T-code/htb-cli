package utils

import (
	"io"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
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

func GetMachineType(machine_id interface{}) string {
	machine_id = fmt.Sprintf("%v", machine_id)
	url := "https://www.hackthebox.com/api/v4/machine/profile/" + machine_id.(string)
	resp := HtbGet(url)
	info := ParseJsonMessage(resp, "info").(map[string]interface{})
	if info["active"].(float64) == 1 {
		return "active"
	} else if info["retired"].(float64) == 1 {
		return "retired"
	}
	return "release"
}

func GetActiveMachineID() interface{} {
	url := "https://www.hackthebox.com/api/v4/machine/active"
	resp := HtbGet(url)
	info := ParseJsonMessage(resp, "info")
	if (info == nil) {
		url = "https://www.hackthebox.com/api/v4/release_arena/active"
		resp = HtbGet(url)
		info = ParseJsonMessage(resp, "info")
		return info.(map[string]interface{})["id"]
	} else {
		return info.(map[string]interface{})["id"]
	}
	return ""
}