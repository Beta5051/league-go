package lcu

import (
	"errors"
	"github.com/shirou/gopsutil/v3/process"
	"strings"
)

func getProcess(name string) (*process.Process, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	for _, p := range processes {
		if n, _ := p.Name(); name == n {
			return p, nil
		}
	}

	return nil, errors.New("process not found")
}

func getCommandArgs(process *process.Process) (map[string]string, error) {
	cmdline, err := process.CmdlineSlice()
	if err != nil {
		return nil, err
	}

	args := map[string]string{}
	for _, line := range cmdline {
		if len(line) > 0 && strings.Contains(line, "=") {
			items := strings.Split(line, "=")
			args[items[0]] = items[1]
		}
	}

	return args, nil
}
