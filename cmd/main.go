package main

import (
	"fmt"
	"github.com/lianzhilu/tiktokmall/cmd/automigrate"
	"github.com/lianzhilu/tiktokmall/cmd/gormgen"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "ttmall",
		Short: "ttmall",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("please enter an option")
		},
	}
	cmd.AddCommand(automigrate.NewCommand())
	cmd.AddCommand(gormgen.NewCommand())
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
