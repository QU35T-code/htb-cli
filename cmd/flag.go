package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
	"log"
	"strconv"
)

var flagCmd = &cobra.Command{
	Use:   "flag",
	Short: "Submit a flag (user and root)",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://www.hackthebox.com/api/v4/machine/own"
		machine_id := utils.GetConfigValue("machineid")
		flag := args[0]
		difficulty, err := strconv.Atoi(args[1])
		if err != nil {
			log.Println(err)
		}
		difficulty2 := strconv.Itoa(difficulty * 10)
		var jsonData = []byte(`{"flag": "` + flag + `", "id": ` + machine_id + `, "difficulty": ` + difficulty2 + `}`)
		resp := utils.HtbPost(url, jsonData)
		message := utils.ParseJsonMessage(resp, "message")
		fmt.Println(message)
	},
}

func init() {
	rootCmd.AddCommand(flagCmd)
}
