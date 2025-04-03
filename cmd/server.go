package cmd

import (
	"github.com/bugbundle/phantom/api"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().String("addr", "127.0.0.1:8080", "Listening interface")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run phantom in server mode",
	Run: func(cmd *cobra.Command, args []string) {
		addr, err := cmd.Flags().GetString("addr")
		if err != nil {
			log.Fatal("Can't retrieve addr flag.")
		}
		api.Server(addr)
	},
}
