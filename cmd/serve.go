package cmd

import (
	"log"
	"net/http"

	"github.com/belarte/metrix2/router"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting web server on :8080...")
		log.Fatal(http.ListenAndServe(":8080", router.New()))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
