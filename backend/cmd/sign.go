/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"orbit_nft/cmd/sign"

	"github.com/spf13/cobra"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Signing services",
	Long:  `Signing services.`,
}

func init() {
	rootCmd.AddCommand(signCmd)
	signCmd.AddCommand(sign.KeygenCmd)
	signCmd.AddCommand(sign.ServiceCmd)
}
