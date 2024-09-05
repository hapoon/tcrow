package main

import (
	"fmt"
	"os"

	"github.com/hapoon/tcrow/action"
	"github.com/urfave/cli/v2"
)

const (
	name    = "tcrow"
	version = "0.1.0"
)

func main() {
	app := &cli.App{
		Version:              version,
		Description:          "A CLI application for TimeCrowd",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "V", Usage: "verbose"},
		},
		Commands: []*cli.Command{
			{
				Name:      "init",
				Usage:     "Initialize setting",
				UsageText: "Initialize setting",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "token", Usage: "アクセストークン"},
				},
				Action: action.Init,
			},
			{
				Name:      "user",
				Aliases:   []string{"u"},
				Usage:     "User",
				UsageText: "User",
				Subcommands: []*cli.Command{
					{
						Name:   "get",
						Usage:  "Get user resource",
						Action: action.GetUser,
					},
					{
						Name:   "get_info",
						Usage:  "Get user info",
						Action: action.GetUserInfo,
					},
					{
						Name:   "get_working",
						Usage:  "Get user working status",
						Action: action.GetUserWorking,
					},
				},
			},
			{
				Name:      "time_entry",
				Aliases:   []string{"te"},
				Usage:     "Time entry",
				UsageText: "Time entry",
				Subcommands: []*cli.Command{
					{
						Name:   "start",
						Usage:  "Create a time entry",
						Action: action.CreateTimeEntry,
					},
					{
						Name:   "stop",
						Usage:  "Stop a time entry",
						Action: action.StopTimeEntry,
					},
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
