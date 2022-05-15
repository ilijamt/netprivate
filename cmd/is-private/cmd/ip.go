package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
)

// ipCmd represents the base command when called without any subcommands
var ipCmd = &cobra.Command{
	Use:   "ip <ip-address>",
	Short: "Is an IP address private IPv4/IPv6?",
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if nil == net.ParseIP(args[0]) {
			return fmt.Errorf("invalid IP address")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return checkIp(net.ParseIP(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}
