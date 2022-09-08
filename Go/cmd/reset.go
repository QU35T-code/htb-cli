package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
	"encoding/json"
	"io"
	"log"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset a machine",
	Long: "Reset a machine",
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://www.hackthebox.com/api/v4/vm/reset"
		machine_id := "492" // Temp
		var jsonData = []byte(`{"machine_id": ` + machine_id + `}`)
		resp := utils.HtbPost(url, jsonData)
		json_body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var result map[string]interface{}
		json.Unmarshal([]byte(json_body), &result)
    	fmt.Println(result["message"])
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
