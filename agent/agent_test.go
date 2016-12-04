package agent

import (
	"fmt"
	"sync"
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

var wg sync.WaitGroup

func TestApplyCommands(t *testing.T) {

	worker := CommandWorker{
		make(chan CommandRequest, 10),
		make(chan CommandResponse),
	}

	// the commands to execute.
	cmds := []string{"sleep 0.1", "sleep 0.05", "sleep 0.2", "sleep 0.05", "sleep 0.15", "sleep 0.05"}

	// apply bounding around number of allowed, concurrent UNIX commands.
	numWorkers := 3
	wg.Add(numWorkers + 2)
	for n := 0; n < numWorkers; n++ {
		go worker.Work(&wg)
	}

	// send commands into the request channel.
	go func() {
		for _, cmd := range cmds {
			fmt.Println("sending command", cmd)
			worker.ReqC <- CommandRequest{time.Now(), cmd}
		}
		close(worker.ReqC)
		wg.Done()
	}()

	// receive responses from response channel.
	go func() {
		for i := 1; i <= len(cmds); i++ {
			response := <-worker.RspC
			fmt.Printf("received response: %d bytes after %v duration with %v\n",
				len(response.Output), response.Elapsed, response.Error)
		}
		wg.Done()
	}()

	wg.Wait()
}
