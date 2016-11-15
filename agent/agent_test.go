package agent

import (
	"fmt"
	"testing"
	"time"
)

func TestSomeFuncOnAgent(t *testing.T) {
	agent := NewAgentLocal()
	//agent := NewAgent("localhost:1979")
	//agent := NewAgentTLS("localhost:1979","cacert.crt","cacert.key")
	fmt.Println("TODO: choose sequence of random ports and create in-process server/client?")
	agent.SomeFuncOnAgent()
}

func TestResponse(t *testing.T) {
	command := CommandRequest{time.Now(), "date"}
	response := CommandResponse{
		command,
		[]byte("Mon Nov 14 22:23:10 CST 2016"),
		time.Now(),
		nil,
		0}
	fmt.Println("duration:", response.Elapsed)
}

func TestApplyCommand(t *testing.T) {
	command := CommandRequest{time.Now(), "/bin/true"}
	//command := CommandRequest{time.Now(), "/bin/false"}

	ch := make(chan CommandResponse)
	go ApplyCommand(command, ch)
	response := <-ch

	fmt.Printf("%d bytes applied after %v duration with %v\n",
		len(response.Output), response.Elapsed, response.Error)
}

func TestApplyCommands(t *testing.T) {
	cmds := []string{"sleep 0.2", "sleep 0.1", "sleep 0.3"}

	creq := make(chan CommandRequest)
	crsp := make(chan CommandResponse)

	go CommandWorkQueue(creq, crsp)

	// send commands into the request channel
	for _, cmd := range cmds {
		creq <- CommandRequest{time.Now(), cmd}
	}
	for i := 1; i <= len(cmds); i++ {
		response := <-crsp
		fmt.Printf("just received %d bytes after %v duration with %v\n",
			len(response.Output), response.Elapsed, response.Error)
	}
}
