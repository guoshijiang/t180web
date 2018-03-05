package main

import (
	"t180web/pkg/cmd"
	"os"
)

func main(){
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}












