// Here is the equivalent implementation of the Proxy design pattern in Go based on the provided Java code:

// Copy
package main

import (
	"fmt"
	"os/exec"
)

// CommandExecutor interface
type CommandExecutor interface {
	RunCommand(cmd string) error
}

// CommandExecutorImpl struct
type CommandExecutorImpl struct{}

func (ce *CommandExecutorImpl) RunCommand(cmd string) error {
	// some heavy implementation
	command := exec.Command("sh", "-c", cmd)
	err := command.Run()
	if err != nil {
		return fmt.Errorf("'%s' command execution failed: %v", cmd, err)
	}
	fmt.Printf("'%s' command executed.\n", cmd)
	return nil
}

// CommandExecutorProxy struct
type CommandExecutorProxy struct {
	isAdmin   bool
	executor  CommandExecutor
}

func NewCommandExecutorProxy(user, pwd string) *CommandExecutorProxy {
	isAdmin := user == "Pankaj" && pwd == "J@urnalD$v"
	return &CommandExecutorProxy{
		isAdmin:  isAdmin,
		executor: &CommandExecutorImpl{},
	}
}

func (cep *CommandExecutorProxy) RunCommand(cmd string) error {
	if cep.isAdmin {
		return cep.executor.RunCommand(cmd)
	}

	if cmdTrimmed := trimCommand(cmd); cmdTrimmed == "rm" {
		return fmt.Errorf("rm command is not allowed for non-admin users.")
	}

	return cep.executor.RunCommand(cmd)
}

func trimCommand(cmd string) string {
	return cmd[:2]
}

func main() {
	executor := NewCommandExecutorProxy("Pankaj", "wrong_pwd")

	err := executor.RunCommand("ls -ltr")
	if err != nil {
		fmt.Println("Exception Message:", err)
	}

	err = executor.RunCommand("rm -rf abc.pdf")
	if err != nil {
		fmt.Println("Exception Message:", err)
	}
}
// In this Go implementation:

// CommandExecutor is an interface specifying the RunCommand method.
// CommandExecutorImpl is a struct implementing the CommandExecutor interface with the actual heavy implementation.
// CommandExecutorProxy is a struct acting as a proxy, controlling access based on the user credentials and restricting certain commands.
// Note: The error handling in Go is used to handle exceptions similar to how exceptions are handled in Java. The exec package is used to run commands, and the error returned by the RunCommand method contains information about the execution result.

// Please adjust the code as needed based on your specific requirements and preferences.