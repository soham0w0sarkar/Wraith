package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "wraith",
	Short: "Wraith.shh - Share isolated sandboxes with ease",
	Long: `Wraith is a CLI tool that spins up an isolated dev environment
and makes it shareable over the web.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Wraith CLI!")
		fmt.Println("Use 'wraith serve' to start the local server.")
		fmt.Println("Use 'wraith share' to share your environment.")
		fmt.Println("Use 'wraith stop' to stop the server.")
		fmt.Println("Use one of the subcommands: serve, share, stop")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
