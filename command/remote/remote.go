package remote

import "github.com/mitchellh/cli"

func Factory() (cli.Command, error) {
	return &cmd{}, nil
}

type cmd struct {

}

func (c *cmd) Help() string {
	panic("implement me")
}

func (c *cmd) Run(args []string) int {
	panic("implement me")
}

func (c *cmd) Synopsis() string {
	panic("implement me")
}
