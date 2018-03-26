package main

import (
	"fmt"

	"github.com/urfave/cli"
	"os"
)

func main() {
	fmt.Printf("ok\n")
	var lang string
	app := cli.NewApp()
	app.Name = "boom"
	app.Version = "0.1"
	app.Usage = "make easy"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "l",
			Value:       "e",
			Usage:       "greet for lang",
			Destination: &lang,
		},
	}
	app.Commands = []cli.Command{
		start,
	}
	app.Before = func(c *cli.Context) error {
		fmt.Printf("before func args is %v\n", c.Args())
		return nil
	}
	app.Action = func(c *cli.Context) {
		fmt.Printf("in action len is %d,args is %v\n", c.NArg(), c.Args())
		//fmt.Printf("greet first context is %v\n", c.Args().First())
		fmt.Printf("lang is %v\n", lang)

	}
	app.Run(os.Args)
}
