package api

import (
	"context"
	"log"
	"net/http"
	"orbit_nft/util"
	"os"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// StaticCmd represents the static command
var StaticCmd = &cobra.Command{
	Use:   "static",
	Short: "Run a (quick and dirty) static file server",
	Long:  `Run a static file server, quick and dirty. No security guaranteed. Use only for **development**.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			panic(err)
		}
		static(":" + port)
	},
}

func init() {
	StaticCmd.Flags().StringP("port", "p", "8081", "Port opened for the server")
}

func static(port string) {
	router := gin.Default()
	router.Static("/assets/nft", os.Getenv("LOCAL_STORAGE_DIR"))
	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	util.RegisterOSSignalHandler(func() {
		log.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}
	}, syscall.SIGINT, syscall.SIGTERM)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("listen: %s\n", err)
	}
}
