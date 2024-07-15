package defaults

//import "os"

const (
  ServerPort uint16 = 5000
  ServerTickets string = "./store/tickets"
  ServerUsers string = "./store/users/users.json"
  ServerEmails string = "./store/emails"

  LogVerbose bool = false
  LogLevel string = "info"

  CliHost string = "localhost"
  CliPort uint16 = 5000
)

// exit code struct using iota : creates a sequence of incrementing numbers in const dec
type ExitCode int

const (
  ExitSuccess ExitCode = iota
  ExitStartError
  ExitShutdownError
)
