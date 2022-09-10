package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
	"encoding/json"
	"io"
	"log"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current machine",
	Long: "Stop the current machine",
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := utils.GetConfigValue("machineid")
		url := "https://www.hackthebox.com/api/v4/machine/stop"
		var jsonData = []byte(`{"machine_id": ` + machine_id + `}`)
		resp := utils.HtbPost(url, jsonData)
		json_body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var result map[string]interface{}
		json.Unmarshal([]byte(json_body), &result)
		fmt.Println("Active Machines")
    	fmt.Println(result["message"])

		// Retired
		url = "https://www.hackthebox.com/api/v4/vm/terminate"
		var jsonData2 = []byte(`{"machine_id": ` + machine_id + `}`)
		resp = utils.HtbPost(url, jsonData2)
		json_body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var result2 map[string]interface{}
		json.Unmarshal([]byte(json_body), &result2)
		fmt.Println("Retired Machines")
    	fmt.Println(result2["message"])
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
