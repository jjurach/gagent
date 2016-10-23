
# Overview

Start an interactive command from a UNIX shell which allows you to target (a
potentially large set of) machines by sending commands and responding to
output. For example, run the "date" command on all machines of a cluster to
eyeball whether NTP is likely running.

# vs. Competitors

Chef and Ansible are heavyweight with bulky syntax ruby and python footprints.
They solve higher level problems.

Salt and Puppet are declarative, also solving much more complicated problems.

# Warning

Misuse of this module may be highly insecure and destructive. Please use a
measure of caution who can connect to your daemon and what necessary access to
remote systems your daemon may expose when you enable insecure settings.

# What is Involved

Execute remote agent (after possible replication). Remote agent subscribes to
request messages. Upon request, run "bash -c 'command'" and capture its
stdout, stderr, exit status, timing info, etc. into a message and publish this
response somewhere. Remote agent also handles: ping, read file, write file, and
info.

A daemon can send and receive messages with agents. A daemon may use SSH to
start and manage multiple remote agents, and communicate with STDIN and STDOUT
of those connections to transmit messages. A daemon may also listen (e.g. TLS
on port 9007) for connections from other daemons, or console.

Readline-based interactive console accepts commands from user, sends command
requests to remote agents through a daemon.  Console attempts connection to
local daemon, and on failure, starts a local daemon, and re-attempts connection
to it. Other daemon clients may generate requests and process responses. Daemon
may remain resident after the daemon client exits.

When daemon receives a request for a machine for which it cannot already send
and receive messages, the daemon uses SSH to do the following activities: test
for existence of agent installation, install agent, execute remote agent,
send and receive command requests and responses with that remote agent.

The console binary is a single file capable of handling all the roles above. As
such, this binary can be copied to the remote machine before being invoked with
arguments to apply the remote agent behavior. 

Command requests have unique IDs, may be sent asynchronously, so responses may
arrive out of order. Agent is expected to respond immediately to pings, so
daemon sends regular heartbeats, and can use SSH to verify whether the remote
process is running, and can kill one or more remote sub-processes to recover
control of agent or can kill the agent itself.

Agent replies with the PID of sub-process immediately after starting the
command. Then after the sub-process exits, the agent sends another message
containing the response. Agent may refuse command requests if its command queue
exceeds a maximum threshold. Each command executes in its own directory, which
is typically removed upon successful communication of the response.

A config file and a private state file can be used to define tags related to
machines and to keep track of whether or where the remote replication exists.
The config file can decide whether to connect to a resident daemon, and for
shared accounts, whether to require user identify herself upon console startup.

# Concepts

  agent
  agent/handler
  handler
  handler/agent
  process
  config
  state
  queue
  daemon
  console

# Use Cases

  console interacts with user
  console attempts message to daemon
  console creates daemon process
  console sends and receives messages to daemon
  daemon attempts message to agent
  daemon replicates to remote machine
  daemon creates agent process
  daemon sends and receives messages to agent
  daemon logs consoles, commands, and responses
  agent replies to ping
  agent replies to info
  agent reads or writes a file
  agent runs a command, captures and transmits output
  queue notifies listeners of availabilty of incoming message
  queue allows atomic, thread-safe pop() for multiple threads

