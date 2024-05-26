package cmd

import (
	_ "embed"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

//go:generate sh -c "printf %s $(git rev-parse --short HEAD) > commit.txt"
//go:embed commit.txt
var Commit string

//go:generate sh -c "cat ../version.md > version.txt"
//go:embed version.txt
var Version string

var debug bool

var rootCmd = &cobra.Command{
	Use:   "ignition",
	Short: "A unix-like Init system",
	Long:  "Ignition is a unix-like Init system",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Ignition",
	Run: func(cmd *cobra.Command, args []string) {
		println("Ignition version", Version)
		println("Commit: ", Commit)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
}

func preRunHook(cmd *cobra.Command, args []string) {
	setupLogger()
}

func setupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func Iginite() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
