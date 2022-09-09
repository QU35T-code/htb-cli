package cmd

import (
	"github.com/spf13/cobra"
	"htb/utils"
	"strconv"
	"log"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		// machine_id := utils.SearchMachineIDByName(args[0])

		// Temp
		machine_id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Println(err)
		}
		machine_id2 := strconv.Itoa(machine_id)
		utils.SetConfigValue("machineID", machine_id2)

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
