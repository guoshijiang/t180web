package cmd

import (
	"github.com/spf13/cobra"
	"t180web/pkg/cli"
)

var RootCmd = &cobra.Command{
	Use: "t180webServer",
	Short: "t180webServer App cmd",
	Long: `t180webServer is a tool to manage golang servic`,
	Example:`
		t180webServer t180
	`,
}

func init(){
	RootCmd.AddCommand(cli.HeartDataCmd("t180"))
}