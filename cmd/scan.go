package cmd

import (
	log "github.com/mdeheij/logwrap"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan library for music files",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("Scanning not implemented yet")
	},
}

func init() {
	RootCmd.AddCommand(scanCmd)
}
