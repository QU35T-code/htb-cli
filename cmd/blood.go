package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
	"os"
)

var bloodCmd = &cobra.Command{
	Use:   "blood",
	Short: "Displays users who have blood the machine",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := utils.GetConfigValue("machineid")
		url := "https://www.hackthebox.com/api/v4/machine/profile/" + machine_id
		resp := utils.HtbGet(url)
		info := utils.ParseJsonMessage(resp, "info").(map[string]interface{})
		if (info["user_owns_count"].(float64) == 0) {
			fmt.Println("There is no first blood for the user")
		}
    	if (info["root_owns_count"].(float64) == 0) {
			fmt.Println("There is no first blood for the root")
		}
		if (info["root_owns_count"].(float64) == 0 || info["user_owns_count"].(float64) == 0) {
			os.Exit(1)
		}
		infoUserBlood := info["userBlood"].(map[string]interface{})["user"].(map[string]interface{})
		infoRootBlood := info["rootBlood"].(map[string]interface{})["user"].(map[string]interface{})
		fmt.Print("Machine : ")
    	fmt.Println(info["name"])
    	fmt.Println("")
    	fmt.Println("--- USER ---")
    	fmt.Print("Name : ")
		fmt.Println(infoUserBlood["name"])
    	fmt.Print("Time : ")
    	fmt.Println(info["firstUserBloodTime"])
		fmt.Println("")
		fmt.Println("--- Root ---")
		fmt.Print("Name : ")
    	fmt.Println(infoRootBlood["name"])
		fmt.Print("Time : ")
    	fmt.Println(info["firstRootBloodTime"])
	},
}

func init() {
	rootCmd.AddCommand(bloodCmd)
}
