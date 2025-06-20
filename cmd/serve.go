package cmd

import (
	"log"
	"net/http"
	"github.com/spf13/cobra"
	"github.com/belarte/metrix2/web/templates"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server",
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			templates.HomePage().Render(r.Context(), w)
		})
		log.Println("Starting web server on :8080...")
		log.Fatal(http.ListenAndServe(":8080", nil))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
