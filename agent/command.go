package agent

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type CommandRequest struct {
	TimeStamp time.Time
	Cmd       string
}

type CommandResponse struct {
	Request   CommandRequest
	Output    []byte
	StartTime time.Time
	Error     error
	Elapsed   time.Duration
}

func (r CommandRequest) String() string {
	return fmt.Sprintf("Command{%v, [[%s]]}", r.TimeStamp, r.Cmd)
}

func (r CommandResponse) String() string {
	return fmt.Sprintf("CommandResponse{%v, [[%s]], %v, %v, %v}", r.Request,
		r.Output, r.StartTime, r.Error, r.Elapsed)
}

func ApplyCommands(commands []string, ch chan<- CommandResponse) {
	for _, cmd := range commands {
		command := CommandRequest{
			time.Now(),
			cmd,
		}
		fmt.Println("applying", command)
		ApplyCommand(command, ch)
	}
}

func ApplyCommand(command CommandRequest, ch chan<- CommandResponse) {
	startTime := time.Now()
	//fmt.Println("applying", command)
	fields := strings.Fields(command.Cmd)
	output, err := exec.Command(fields[0], fields[1:]...).Output()
	elapsed := time.Now().Sub(startTime)
	response := CommandResponse{
		Request:   command,
		StartTime: startTime,
		Error:     err,
		Output:    output,
		Elapsed:   elapsed,
	}
	//fmt.Println("applied", response)
	ch <- response
}
