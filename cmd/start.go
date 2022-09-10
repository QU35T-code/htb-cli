package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a machine",
	Long: `Starts a Hackthebox machine specified in argument`,
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := utils.GetConfigValue("machineid")
		machine_type := utils.GetMachineType(machine_id)
		if machine_type == "active" {
			url := "https://www.hackthebox.com/api/v4/machine/play/" + machine_id
			var jsonData = []byte("{}")
			resp := utils.HtbPost(url, jsonData)
			message := utils.ParseJsonMessage(resp, "message")
			fmt.Print("Active : ")
			fmt.Println(message)
		} else if machine_type == "retired" {
			url := "https://www.hackthebox.com/api/v4/vm/spawn"
			var jsonData2 = []byte(`{"machine_id": ` + machine_id + `}`)
			resp := utils.HtbPost(url, jsonData2)
			message := utils.ParseJsonMessage(resp, "message")
			fmt.Print("Retired : ")
			fmt.Println(message)
		} else {
			url := "https://www.hackthebox.com/api/v4/release_arena/spawn"
			var jsonData3 = []byte(`{}`)
			resp := utils.HtbPost(url, jsonData3)
			message := utils.ParseJsonMessage(resp, "message")
			fmt.Print("Release : ")
			fmt.Println(message)
		}		
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
