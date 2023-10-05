package remote

import (
	"os"
	"path/filepath"
	"strings"
)

func RunKeyScan(host string) error {
	sshDir := filepath.Join(os.Getenv("HOME"), ".ssh")
	knownHostsFile := filepath.Join(sshDir, "known_hosts")
	keys, err := scanKeys(host)
	if err != nil {
		return err
	}
	b, err := os.ReadFile(knownHostsFile)
	if err != nil {
		return err
	}
	lines := strings.Split(string(b), "\n")
	for _, key := range keys {
		found := false
		for _, line := range lines {
			if line == key {
				found = true
				break
			}
		}
		if !found {
			lines = append(lines, key)
		}
	}
	b = []byte{}
	for _, line := range lines {
		b = append(b, []byte(line+"\n")...)
	}
	return os.WriteFile(knownHostsFile, b, os.ModePerm)
}
