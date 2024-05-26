package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:    "config",
	Short:  "Manage Ignition configuration",
	Long:   "Manage Ignition configuration blobs",
	PreRun: preRunHook,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("config command")
		log.Debug().Msg("config command debug")
	},
}
