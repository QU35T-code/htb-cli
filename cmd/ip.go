package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
	"os"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get machine IP address",
	Long: `Retrieves the IP address of the machine deployed on Hackthebox and displays it`,
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := utils.GetConfigValue("machineid")
		url := "https://www.hackthebox.com/api/v4/machine/profile/" + machine_id
		resp := utils.HtbGet(url)
		info := utils.ParseJsonMessage(resp, "info")
		infomap := info.(map[string]interface{})
    	if infomap["ip"] == nil {
			fmt.Println("Machine is down")
			os.Exit(1)
		}
		fmt.Println(infomap["name"])
		fmt.Println(infomap["ip"])
	},
}



func init() {
	rootCmd.AddCommand(ipCmd)
}
