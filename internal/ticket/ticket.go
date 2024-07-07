package ticket

import (
  "flag"
	"github.com/sfaizh/ticket-management-system/internal/structs/defaults"
)

var (
  port = flag.Uint("port", uint(defaults.ServerPort), "server `port`")
  tickets = flag.String("tickets", defaults.ServerTickets, "tickets `directory`")
  users = flag.String("users", defaults.ServerUsers, "users `file`")
  emails = flag.String("mails", defaults.ServerEmails, "mails `directory")
  verbose = flag.Bool("verbose", defaults.LogVerbose, "verbose logs")
  logLevel = flag.String("logLevel", defaults.LogLevel, "logs verbosity either 'info' | 'debug'")
)

func PrintPtr(iptr int) *int {
	return &iptr
}

 
