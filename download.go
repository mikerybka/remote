package remote

import (
	"os"
	"path/filepath"

	"github.com/mikerybka/util"
)

// Download downloads a file from the internet and copies it to the machine.
func Download(user, host, dst, url string) error {
	path := filepath.Join(os.TempDir(), "remote", util.RandomID())
	err := util.Download(url, path)
	if err != nil {
		return err
	}
	return Copy(user, host, dst, path)
}
