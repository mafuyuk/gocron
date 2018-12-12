package main

import (
	"fmt"
	"io"
)

const displayHeaderFormat = "%-7s %-7s %s"
const displayFormat = "%-7d %-7d %s"

var jobStringHeader = fmt.Sprintf(displayHeaderFormat, "PID", "Status", "Label")

type Jobs []*Job

type Service struct {
	Jobs              Jobs
	CurrentUserFilter []Label
}

func (s Service) display(out io.Writer) (err error) {
	_, err = fmt.Fprintln(out, jobStringHeader)
	for _, job := range s.Jobs {
		_, err = fmt.Fprintln(out, job)
	}
	return err
}

func (s Service) displayForCurrentUserFilter(out io.Writer) (err error) {
	_, err = fmt.Fprintln(out, jobStringHeader)
	for _, job := range s.Jobs {
		for _, filter := range s.CurrentUserFilter {
			if filter == job.Label {
				_, err = fmt.Fprintln(out, job)
			}
		}
	}
	return err
}

type Label string

type Job struct {
	PID    int
	Status int
	Label
}

func (j *Job) String() string {
	return fmt.Sprintf(displayFormat, j.PID, j.Status, j.Label)
}
