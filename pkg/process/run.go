package process

import (
	"os"
	"os/exec"
)

type Runnable struct {
	Name string
	Args []string
	Wait bool
}

func (r Runnable) run() (cmd *exec.Cmd, err error) {
	if r.Wait {
		return RunWithWait(r.Name, r.Args...)
	}
	return RunWithoutWait(r.Name, r.Args...)
}

func RunWithWait(name string, args ...string) (cmd *exec.Cmd, err error) {
	cmd = exec.Command(name, args...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	err = cmd.Start()
	if err != nil {
		return
	}
	err = cmd.Wait()
	return
}

func RunWithoutWait(name string, args ...string) (cmd *exec.Cmd, err error) {
	cmd = exec.Command(name, args...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	err = cmd.Start()
	return
}
