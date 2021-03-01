package cmd

import (
	"github.com/inexio/thola/core/request"
	"github.com/spf13/cobra"
)

func init() {
	addDeviceFlags(identifyCMD)
	rootCMD.AddCommand(identifyCMD)
}

var identifyCMD = &cobra.Command{
	Use:   "identify [host]",
	Short: "Automatically identify devices",
	Long: "Automatically identify devices.\n\n" +
		"It returns properties like vendor, model, serial number,...",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		r := request.IdentifyRequest{
			BaseRequest: getBaseRequest(args[0]),
		}
		handleRequest(&r)
	},
}
