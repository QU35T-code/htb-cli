package cmd

import (
	"fmt"
	"htb-cli/utils"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current machine",
	Long:  "Stop the current machine",
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := utils.GetActiveMachineID()
		machine_type := utils.GetMachineType(machine_id)
		machine_id = fmt.Sprintf("%v", machine_id)
		if machine_type == "active" {
			url := "https://www.hackthebox.com/api/v4/machine/stop"
			var jsonData = []byte(`{"machine_id": ` + machine_id.(string) + `}`)
			resp := utils.HtbPost(url, jsonData)
			message := utils.ParseJsonMessage(resp, "message")
			fmt.Print("Active : ")
			fmt.Println(message)
		} else if machine_type == "retired" {
			url := "https://www.hackthebox.com/api/v4/vm/terminate"
			var jsonData2 = []byte(`{"machine_id": ` + machine_id.(string) + `}`)
			resp := utils.HtbPost(url, jsonData2)
			message := utils.ParseJsonMessage(resp, "message")
			fmt.Print("Retired : ")
			fmt.Println(message)
		} else {
			url := "https://www.hackthebox.com/api/v4/release_arena/terminate"
			var jsonData2 = []byte(`{}`)
			resp := utils.HtbPost(url, jsonData2)
			message := utils.ParseJsonMessage(resp, "message")
			fmt.Print("Release : ")
			fmt.Println(message)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
