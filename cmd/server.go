package cmd

import (
	"lgmontenegro/sword/internal/service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Up with API router server",
	Long:  `Runs the server for API Router`,
	Run: func(cmd *cobra.Command, args []string) {
    server := service.ConfigServer()
    server.Run()
	},
}
