package main

import (
	"github.com/urfave/cli"
	"fmt"
)

var start=cli.Command{
	Name:"start",
	Usage:"start this",
	Description:"test",
	Action:func(context *cli.Context) error {
		name:=context.Command.Name
		fmt.Printf("start command ok,name is %v\n",name)
		return nil
	},
}