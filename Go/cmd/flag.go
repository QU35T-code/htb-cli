package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
	"encoding/json"
	"io"
	"log"
	"strconv"
)

var flagCmd = &cobra.Command{
	Use:   "flag",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://www.hackthebox.com/api/v4/machine/own"
		machine_id := "492" // Temp
		flag := args[0]
		difficulty, err := strconv.Atoi(args[1])
		if err != nil {
			log.Println(err)
		}
		difficulty2 := strconv.Itoa(difficulty * 10)
		var jsonData = []byte(`{"flag": "` + flag + `", "id": ` + machine_id + `, "difficulty": ` + difficulty2 + `}`)
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
	rootCmd.AddCommand(flagCmd)
}
