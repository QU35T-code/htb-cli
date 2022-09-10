package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
	"io"
	"log"
	"encoding/json"
	"os"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {		
		url := "https://www.hackthebox.com/api/v4/machine/active"
		resp := utils.HtbGet(url)
		json_body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var result map[string]interface{}
		json.Unmarshal([]byte(json_body), &result)
		info := result["info"]
		if (info == nil) {
			fmt.Println("No machine is active")
			os.Exit(1)
		}
		infomap := info.(map[string]interface{})
    	fmt.Print("ID : ")
    	fmt.Println(infomap["id"])
		fmt.Print("Name : ")
    	fmt.Println(infomap["name"])
		fmt.Print("Plan : ")
    	fmt.Println(infomap["type"])
		fmt.Print("Server : ")
    	fmt.Println(infomap["lab_server"])
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
