package cmd

import (
	"fmt"
	"gcp-localstack/core"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gcp-localstack",
	Short: "A local emulator for GCP services",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the local emulator",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting GCP LocalStack...")
		for _, plugin := range core.ListPlugins() {
			if err := core.StartPlugin(plugin); err != nil {
				fmt.Printf("Failed to start plugin %s: %v\n", plugin, err)
			} else {
				fmt.Printf("Started plugin %s\n", plugin)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
