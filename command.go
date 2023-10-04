package remote

import (
	"bytes"

	"github.com/mikerybka/util"
	"golang.org/x/crypto/ssh"
)

type Command struct {
	Cmd   string
	Dir   string
	Stdin []byte
}

func (c *Command) Run(on *Machine) ([]byte, error) {
	sshConfig, err := util.SSHClientConfig(on.UserID)
	if err != nil {
		return nil, err
	}

	client, err := ssh.Dial("tcp", on.Host+":22", sshConfig)
	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	session.Stdin = bytes.NewReader([]byte(c.Stdin))

	out, err := session.CombinedOutput(c.Cmd)
	session.Close()
	client.Close()
	return out, err
}
