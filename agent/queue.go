package agent

type CommandWorker struct {
	ID       int
	Work     chan CommandRequest
	QuitChan chan bool
}

type CommandQueue struct {
	WorkQueue chan WorkRequest
}

type WorkRequest struct {
	Request  CommandRequest
	Response chan CommandResponse
}

func CommandWorkQueue(creq chan CommandRequest, crsp chan CommandResponse) {
	for {
		command := <-creq
		go ApplyCommand(command, crsp)
	}
}
