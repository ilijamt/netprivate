package cmd

import (
	"github.com/spf13/cobra"
	"net"
)

// cidrCmd represents the base command when called without any subcommands
var cidrCmd = &cobra.Command{
	Use:   "cidr <cidr-range>",
	Short: "Is an CIDR range address private IPv4/IPv6?",
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		_, _, err = net.ParseCIDR(args[0])
		return err
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ip, _, _ := net.ParseCIDR(args[0])
		return checkIp(ip)
	},
}

func init() {
	rootCmd.AddCommand(cidrCmd)
}
