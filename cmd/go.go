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
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/jjurach/gagent/agent"
	"github.com/jjurach/gagent/daemon"
	"github.com/spf13/cobra"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Interactive DSL for targeting commands on agents",
	Long:  `Provide a featureful command prompt to user to identify a set of agents, and to execute commands and other operations on those remote agents. Await and collate responses from multiple agents. DSL allows output redirection to pipe into unix commands, for example: grep, awk, tee, sort, sed, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		handleInteraction()
	},
}

func init() {
	RootCmd.AddCommand(goCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func usage(w io.Writer) {
	io.WriteString(w, "commands:\n")
	io.WriteString(w, completer.Tree("    "))
}

// Function constructor - constructs new function for listing given directory
func listFiles(path string) func(string) []string {
	return func(line string) []string {
		names := make([]string, 0)
		files, _ := ioutil.ReadDir(path)
		for _, f := range files {
			names = append(names, f.Name())
		}
		return names
	}
}

var completer = readline.NewPrefixCompleter(
	readline.PcItem("mode",
		readline.PcItem("vi"),
		readline.PcItem("emacs"),
	),
	readline.PcItem("login"),
	readline.PcItem("say",
		readline.PcItemDynamic(listFiles("./"),
			readline.PcItem("with",
				readline.PcItem("following"),
				readline.PcItem("items"),
			),
		),
		readline.PcItem("hello"),
		readline.PcItem("bye"),
	),
	readline.PcItem("setprompt"),
	readline.PcItem("setpassword"),
	readline.PcItem("bye"),
	readline.PcItem("help"),
	readline.PcItem("go",
		readline.PcItem("build", readline.PcItem("-o"), readline.PcItem("-v")),
		readline.PcItem("install",
			readline.PcItem("-v"),
			readline.PcItem("-vv"),
			readline.PcItem("-vvv"),
		),
		readline.PcItem("test"),
	),
	readline.PcItem("sleep"),
)

func handleInteraction() {
	l, err := readline.NewEx(&readline.Config{
		//Prompt:          "\033[31m»\033[0m ",
		Prompt:          "handle://> ",
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold: true,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Println("here")
	daemon := daemon.NewDaemonLocal()
	daemon.SomeFuncOnDaemon()

	setPasswordCfg := l.GenPasswordConfig()
	setPasswordCfg.SetListener(func(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
		l.SetPrompt(fmt.Sprintf("Enter password(%v): ", len(line)))
		l.Refresh()
		return nil, 0, false
	})

	log.SetOutput(l.Stderr())
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "mode "):
			switch line[5:] {
			case "vi":
				l.SetVimMode(true)
			case "emacs":
				l.SetVimMode(false)
			default:
				println("invalid mode:", line[5:])
			}
		case line == "mode":
			if l.IsVimMode() {
				println("current mode: vim")
			} else {
				println("current mode: emacs")
			}
		case line == "login":
			pswd, err := l.ReadPassword("please enter your password: ")
			if err != nil {
				break
			}
			println("you enter:", strconv.Quote(string(pswd)))
		case line == "help":
			usage(l.Stderr())
		case line == "setpassword":
			pswd, err := l.ReadPasswordWithConfig(setPasswordCfg)
			if err == nil {
				println("you set:", strconv.Quote(string(pswd)))
			}
		case strings.HasPrefix(line, "run"):
			cmd := strings.TrimSpace(line[3:])
			if cmd == "" {
				log.Println("Usage: run <command>")
				break
			}
			log.Printf("TODO - need to validate/audit command")
			command := agent.CommandRequest{time.Now(), cmd}
			log.Printf("running %v\n", command)

			//command := CommandRequest{time.Now(), "/bin/false"}

			ch := make(chan agent.CommandResponse)
			go agent.ApplyCommand(command, ch)
			response := <-ch

			fmt.Printf("%d bytes applied after %v duration with error %v\n",
				len(response.Output), response.Elapsed, response.Error)
			fmt.Printf("[%s]\n", strings.TrimSpace(string(response.Output)))

		case strings.HasPrefix(line, "setprompt"):
			prompt := line[10:]
			if prompt == "" {
				log.Println("setprompt <prompt>")
				break
			}
			l.SetPrompt(prompt)
		case strings.HasPrefix(line, "say"):
			line := strings.TrimSpace(line[3:])
			if len(line) == 0 {
				log.Println("say what?")
				break
			}
			go func() {
				for range time.Tick(time.Second) {
					log.Println(line)
				}
			}()
		case line == "bye" || line == "exit" || line == "quit" || line == "q":
			goto exit
		case line == "sleep":
			log.Println("sleep 4 second")
			time.Sleep(4 * time.Second)
		case line == "":
		default:
			log.Println("you said:", strconv.Quote(line))
		}
	}
exit:
}
