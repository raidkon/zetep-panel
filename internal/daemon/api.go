package daemon

// RunRequest is POST /v1/run — execute a z-panel subcommand (argv without program name).
type RunRequest struct {
	Args []string `json:"args"`
}

// RunResponse is the result of a delegated run (stdout/stderr captured from the child process).
type RunResponse struct {
	ExitCode int    `json:"exit_code"`
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
}
