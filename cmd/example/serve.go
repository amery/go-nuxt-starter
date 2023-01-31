package main

import (
	"net/http"

	"github.com/amery/go-nuxt-starter/web/client"
	"github.com/spf13/cobra"
	"go.sancus.dev/web"
)

func NewRouter() http.Handler {
	return http.NotFoundHandler()
}

// Command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves web interface",
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := http.Server{
			Addr: ":8080",
		}

		// application router
		h := NewRouter()

		// middleware
		for _, middleware := range []web.MiddlewareHandlerFunc{
			client.Middleware(),
		} {
			h = middleware(h)
		}

		srv.Handler = h

		return srv.ListenAndServe()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
