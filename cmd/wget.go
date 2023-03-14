package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var help string
var wget = &cobra.Command{
	Use:     "wget",
	Example: "example",
	Short:   "short",
	Long:    "long",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dqdqw")

	},
}

func init() {
	cli.AddCommand(wget)
	wget.Flags().StringVarP(&help, "help", "h", "", "please use -h")
}
