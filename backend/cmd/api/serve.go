package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"orbit_nft/api"
	"orbit_nft/cmd/common"
	"orbit_nft/util"
	"path"
	"syscall"

	"github.com/spf13/cobra"
)

// ServeCmd represents the serve command
var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run an instance of the API server (in the development mode)",
	Long:  `Run an instance of the API server (in the development mode)`,
	Run: func(cmd *cobra.Command, args []string) {
		port := cmd.Flag("port").Value.String()
		dir := cmd.Flag("cert").Value.String()
		serve(port, dir)
	},
}

func init() {
	ServeCmd.Flags().StringP("port", "p", "3000", "Port opened for the server")
	ServeCmd.Flags().StringP("cert", "c", "./secrets", "Path to the directory containing TLS certificate")
}

func serve(port string, dir string) {
	sqlDB := common.MySQL()

	server := api.NewServer(util.GetEnv(), sqlDB)
	server.OnShutdown(func(ctx context.Context) {
		// ensure all routine is closed
		// ensure all external connection is closed
		_ = sqlDB.Close()
	})

	util.RegisterOSSignalHandler(func() {
		log.Info().Msg("Shutting down server...")
		server.Shutdown()
	}, syscall.SIGTERM, syscall.SIGINT)

	log.Info().Msg("Serving...")
	server.Run(":"+port, "", path.Join(dir, "server.pem"), path.Join(dir, "server.key"))
}
