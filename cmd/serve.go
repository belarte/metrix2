package cmd

import (
	"github.com/belarte/metrix2/router"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server",
	Run: func(cmd *cobra.Command, args []string) {
		router.Serve(":8080")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
