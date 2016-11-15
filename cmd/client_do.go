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

	"github.com/jjurach/gagent/daemon"
	"github.com/spf13/cobra"
)

// client_doCmd represents the client_do command
var client_doCmd = &cobra.Command{
	Use:   "client_do",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("Calling ConnectAndApply")
		addr := "127.0.0.1:10443"
		certFile := "keys/ca.pem"
		keyFile := "keys/ca-key.pem"
		daemon.ConnectAndApply(addr, certFile, keyFile)
	},
}

func init() {
	RootCmd.AddCommand(client_doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// client_doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// client_doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
