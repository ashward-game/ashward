package api

import (
	"log"
	"orbit_nft/testutil"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate self-signed ssl keys/cert with sane defaults",
	Long:  `Generate self-signed ssl keys/cert with sane defaults for our development API server.`,
	Run: func(cmd *cobra.Command, args []string) {
		out, _ := cmd.Flags().GetString("out")
		expireIn, _ := cmd.Flags().GetDuration("expire-in")
		if err := os.MkdirAll(out, os.ModePerm); err != nil {
			log.Println(err)
		}
		if err := testutil.CreateTLSCert(out, expireIn); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	InitCmd.Flags().StringP("out", "o", "./secrets", "Path to the directory for storing generated files")
	InitCmd.Flags().DurationP("expire-in", "e", time.Hour, "Certificate expire duration")
}
