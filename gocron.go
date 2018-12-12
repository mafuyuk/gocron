package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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
	isAll := flag.Bool("a", false, "All")
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
		var jobs []*Job
		outLaunchList, err := exec.Command("launchctl", "list").Output()
		if err != nil {
			fmt.Fprintf(c.OutStream, "fail outLaunchList\n")
			return ExitCodeParseFlagError
		}

		rep := regexp.MustCompile("[\\s]+")
		launchList := strings.Split(string(outLaunchList), "\n")[1:]
		for _, v := range launchList {
			if v == "" {
				continue
			}

			job := rep.Split(v, len(launchList))

			var pid int
			if job[0] == "-" {
				pid = -1
			} else {
				pid, err = strconv.Atoi(job[0])
				if err != nil {
					fmt.Fprintf(c.OutStream, "fail strconv.Atoi(job[0])\n")
					fmt.Fprintf(c.OutStream, "%v\n", err)
					return ExitCodeParseFlagError
				}
			}

			status, err := strconv.Atoi(job[1])
			if err != nil {
				fmt.Fprintf(c.OutStream, "fail strconv.Atoi(job[1])\n")
				return ExitCodeParseFlagError
			}
			jobs = append(jobs, &Job{
				PID:    pid,
				Status: status,
				Label:  Label(job[2]),
			})
		}

		s := Service{Jobs: jobs}

		if !*isAll {
			usr, err := user.Current()
			if err != nil {
				fmt.Fprintf(c.OutStream, "fail user.Current\n")
				return ExitCodeParseFlagError
			}

			outPlist, err := exec.Command("ls", usr.HomeDir+"/Library/LaunchAgents").Output()
			if err != nil {
				fmt.Fprintf(c.OutStream, "fail outPlist\n")
				fmt.Fprintf(c.ErrStream, "%v", err)
				return ExitCodeParseFlagError
			}

			plist := strings.Split(string(outPlist), "\n")
			var targetLabels []Label
			for _, v := range plist {
				ext := filepath.Ext(v)
				if ext != ".plist" {
					continue
				}
				targetLabels = append(targetLabels, Label(v[:len(v)-len(ext)]))
			}

			s.CurrentUserFilter = targetLabels
			s.displayForCurrentUserFilter(c.OutStream)
		} else {
			s.display(c.OutStream)
		}

		return ExitCodeOK
	}

	fmt.Fprint(c.OutStream, "Do not run. Use -h\n")

	return ExitCodeOK
}

func main() {
	cli := &cli{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run())
}
