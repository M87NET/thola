package cmd

import (
	"github.com/inexio/thola/core/request"
	"github.com/spf13/cobra"
)

func init() {
	readCMD.AddCommand(readUPSCMD)
}

var readUPSCMD = &cobra.Command{
	Use:   "ups [host]",
	Short: "Read out UPS information of a device",
	Long:  "Read out UPS information of a device like battery capacity and current usage.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		request := request.ReadUPSRequest{
			ReadRequest: getReadRequest(args[0]),
		}
		handleRequest(&request)
	},
}
