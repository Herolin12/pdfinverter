package main

import (
	"bytes"
	"os/exec"
)

// Executor visits my repositories.
type Executor struct {
	command bytes.Buffer
}

// SetCommand writes a command into the buffer.
func (ex *Executor) SetCommand(cmd string) {
	ex.command.WriteString(cmd)
}

// RunCommand executes the command stored
// in the buffer and thenempties the buffer.
func (ex *Executor) RunCommand() {
	cmd := exec.Command("/bin/sh", "-c", ex.command.String())
	if err := cmd.Run(); err != nil {
		ExitErr(err)
	}
	ex.command.Reset()
}
