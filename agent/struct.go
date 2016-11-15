package agent

import (
	"encoding/json"
	"fmt"
)

type Agent struct {
	Addr     string
	CertFile string
	KeyFile  string
}

func NewAgentLocal() *Agent {
	agent := &Agent{}
	return agent
}

func NewAgent(addr string) *Agent {
	agent := &Agent{
		Addr: addr,
	}
	return agent
}

func NewAgentTLS(addr string, certFile string, keyFile string) *Agent {
	agent := &Agent{
		Addr:     addr,
		CertFile: certFile,
		KeyFile:  keyFile,
	}
	return agent
}

func (a Agent) Init() error {
	return nil
}

func (d Agent) SomeFuncOnAgent() {
	j, _ := json.Marshal(d)
	fmt.Println("some function on agent ", string(j))
}
