// Copyright © 2016 James Jurach <james.jurach@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"

	"github.com/jjurach/gagent/daemon"
	"github.com/spf13/cobra"
)

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Accept authenticated connection to initiate messages to and about agents",
	Long:  `[Possibly] detach from caller, listen on tcp port for incoming, authenticated connections to initiate request messages to and about agents. Relay request responses back to connection. Daemon can attempt to start a remote agent to satisfy the target set for command requests, and can manage a large number of such agents.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("daemon will require authentication")
		certFile := "keys/ca.pem"
		keyFile := "keys/ca-key.pem"
		err := daemon.ListenAndHandle(":10443", certFile, keyFile)
		log.Fatal(err)
	},
}

func init() {
	RootCmd.AddCommand(daemonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
