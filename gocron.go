package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const Version string = "v0.1.0"

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
)

type cli struct {
	OutStream, ErrStream io.Writer
}

func (c *cli) Run() int {
	version := flag.Bool("v", false, "Print version information and quit")
	setCron := flag.Bool("e", false, "Setting cron")
	listCron := flag.Bool("l", false, "List cron")
	flag.Parse()

	if *version {
		fmt.Fprintf(c.OutStream, "version %s\n", Version)
		return ExitCodeOK
	}

	if *setCron {
		fmt.Fprintf(c.OutStream, "setCron\n")
		return ExitCodeOK
	}

	if *listCron {
		fmt.Fprintf(c.OutStream, "listCron\n")
		return ExitCodeOK
	}

	fmt.Fprint(c.OutStream, "Do not run. Use -h\n")

	return ExitCodeOK
}

func main() {
	cli := &cli{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run())
}
