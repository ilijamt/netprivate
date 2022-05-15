package cmd

import (
	"fmt"
	"github.com/ilijamt/netprivate"
	"github.com/spf13/cobra"
	"net"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Short: "Is an IP/CIDR address private IPv4/IPv6?",
	Long:  `A small library to test if an IP address is in the private or public range for both IPv4 and IPv6.`,
}

type config struct {
	useExistStatusCode bool
}

var conf config

func init() {
	rootCmd.Flags().BoolVarP(&conf.useExistStatusCode, "use-exit-status-code", "x", false, "Use exist status code as a response, 0 for true/success, anything else is false")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func checkIp(ip net.IP) error {
	if ip == nil {
		return fmt.Errorf("ip is nil")
	}

	if netprivate.Is(ip) {
		if conf.useExistStatusCode {
			os.Exit(0)
		}
		fmt.Println("yes")
	} else {
		if conf.useExistStatusCode {
			os.Exit(1)
		}
		fmt.Println("no")
	}

	return nil
}
