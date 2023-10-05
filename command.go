package remote

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

type Command struct {
	Cmd   string
	Stdin []byte
}

func (c *Command) Run(on *Machine) ([]byte, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "/root"
	}

	// Read private keys
	path := filepath.Join(home, ".ssh/id_rsa")
	pKey, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(pKey)
	if err != nil {
		return nil, err
	}

	// Set up the SSH client configuration.
	config := &ssh.ClientConfig{
		User: on.UserID,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: knownHostsCallback(),
	}

	return run(on.Host, config, c)
}

func run(host string, config *ssh.ClientConfig, cmd *Command) ([]byte, error) {
	// Connect to the remote server.
	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Create an SSH session.
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	session.Stdin = bytes.NewReader([]byte(cmd.Stdin))

	// Run the provided SSH command.
	return session.CombinedOutput(cmd.Cmd)
}

func isKnownHostsError(err error) bool {
	// Check if the error message contains the known hosts error message.
	fmt.Println(err.Error())
	contains := strings.Contains(err.Error(), "knownhosts: key is unknown")
	fmt.Println(contains)
	return contains
}

func scanKeys(host string) ([]string, error) {
	cmd := exec.Command("ssh-keyscan", host)
	outbuf := bytes.NewBuffer(nil)
	cmd.Stdout = outbuf
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(outbuf.String(), "\n")
	filteredLines := []string{}
	for _, line := range lines {
		if line != "" {
			filteredLines = append(filteredLines, line)
		}
	}
	return filteredLines, nil
}

func knownHostsCallback() ssh.HostKeyCallback {
	knownHostsFile := filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")
	knownHostsCallback, err := knownhosts.New(knownHostsFile)
	if err != nil {
		fmt.Printf("Error loading known_hosts file: %v\n", err)
		knownHostsCallback = ssh.InsecureIgnoreHostKey()
	}
	return knownHostsCallback
}
