package lcu

import (
	"encoding/json"
	"fmt"
	"league/internal"
	"net/http"
	"strconv"
)

type Lcu struct {
	Port   uint
	Token  string
	client *internal.Client
}

func NewLcu(port uint, token string) *Lcu {
	lcu := new(Lcu)
	lcu.Port = port
	lcu.Token = token
	lcu.client = internal.NewClient(
		func() string {
			return fmt.Sprintf("https://127.0.0.1:%d", lcu.Port)
		},
		func(req *http.Request) {
			req.SetBasicAuth("riot", lcu.Token)
		},
		func(resp *http.Response) error {
			out := Error{}
			if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
				return err
			}
			return out
		},
	)
	return lcu
}

func NewLcuFromProcess() (*Lcu, error) {
	process, err := getProcess("LeagueClientUx")
	if err != nil {
		return nil, err
	}

	args, err := getCommandArgs(process)
	if err != nil {
		return nil, err
	}

	port, err := strconv.ParseUint(args["--app-port"], 10, 32)
	if err != nil {
		return nil, err
	}
	token := args["--remoting-auth-token"]

	return NewLcu(uint(port), token), nil
}

func (l *Lcu) Send(method, endpoint string, data any, out any) error {
	return l.client.Send(method, endpoint, data, out)
}
