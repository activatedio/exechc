package exechc

// Runtime contains the main runtime configuration
type Runtime struct {
	// Port is the port number to listen on
	Port int
	// Host is the hostname to listen on
	Host string
	// CheckCmd is the static command line to execute
	CheckCmd string
	// CheckExpression is the CEL expression to run against the Output struct of the result
	CheckExpression string
}

// Checker performs actual check for health
type Checker interface {
	Check() (bool, error)
}

// Server handles stop and starting of server processes
type Server interface {
	Start() error
	Shutdown() error
}
