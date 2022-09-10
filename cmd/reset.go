package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset a machine",
	Long: "Reset a machine",
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := utils.GetConfigValue("machineid")
		machine_type := utils.GetMachineType(machine_id)
		if machine_type == "active" {
			url := "https://www.hackthebox.com/api/v4/vm/reset"
			machine_id := utils.GetConfigValue("machineid")
			var jsonData = []byte(`{"machine_id": ` + machine_id + `}`)
			resp := utils.HtbPost(url, jsonData)
			message := utils.ParseJsonMessage(resp, "message")
			fmt.Println(message)
		} else {
			url := "https://www.hackthebox.com/api/v4/release_arena/reset"
			machine_id := utils.GetConfigValue("machineid")
			var jsonData = []byte(`{"machine_id": ` + machine_id + `}`)
			resp := utils.HtbPost(url, jsonData)
			message := utils.ParseJsonMessage(resp, "message")
			fmt.Println(message)
		}

		
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
