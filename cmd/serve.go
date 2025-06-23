package cmd

import (
	"log"
	"net/http"

	"github.com/belarte/metrix2/handlers"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server",
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/", handlers.Home)
		http.HandleFunc("/metrics", handlers.Metrics)
		log.Println("Starting web server on :8080...")
		log.Fatal(http.ListenAndServe(":8080", nil))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
