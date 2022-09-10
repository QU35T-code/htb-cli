package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
)

var bloodCmd = &cobra.Command{
	Use:   "blood",
	Short: "Displays users who have blood the machine",
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := ""
		if len(args) != 0 {
			machine_id = utils.SearchMachineIDByName(args[0])
			machine_id = fmt.Sprintf("%v", machine_id)
		} else {
			machine_id_interface := utils.GetActiveMachineID()
			machine_id = fmt.Sprintf("%v", machine_id_interface)
		}
		url := "https://www.hackthebox.com/api/v4/machine/profile/" + machine_id
		resp := utils.HtbGet(url)
		info := utils.ParseJsonMessage(resp, "info").(map[string]interface{})
		fmt.Printf("Machine : %v\n\n", info["name"])
		if (info["user_owns_count"].(float64) != 0) {
			infoUserBlood := info["userBlood"].(map[string]interface{})["user"].(map[string]interface{})
			fmt.Printf("--- USER ---\nName : %v\nTime : %v\n\n", infoUserBlood["name"], info["firstUserBloodTime"])
		} else {
			fmt.Println("There is no first blood for the user")
		}
    	if (info["root_owns_count"].(float64) != 0) {
			infoRootBlood := info["rootBlood"].(map[string]interface{})["user"].(map[string]interface{})
			fmt.Printf("--- ROOT ---\nName : %v\nTime : %v", infoRootBlood["name"], info["firstRootBloodTime"])
		} else {
			fmt.Println("There is no first blood for the root")
		}
	},
}

func init() {
	rootCmd.AddCommand(bloodCmd)
}
