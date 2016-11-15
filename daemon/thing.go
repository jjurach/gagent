package daemon

import (
	"encoding/json"
	"fmt"
)

type Daemon struct {
	Addr     string
	CertFile string
	KeyFile  string
}

func NewDaemonLocal() *Daemon {
	daemon := &Daemon{}
	return daemon
}

func NewDaemon(addr string) *Daemon {
	daemon := &Daemon{
		Addr: addr,
	}
	return daemon
}

func NewDaemonTLS(addr string, certFile string, keyFile string) *Daemon {
	daemon := &Daemon{
		Addr:     addr,
		CertFile: certFile,
		KeyFile:  keyFile,
	}
	return daemon
}

func (d Daemon) SomeFuncOnDaemon() {
	j, _ := json.Marshal(d)
	fmt.Println("some function on daemon ", string(j))
}
