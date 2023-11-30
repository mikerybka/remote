package remote

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mikerybka/util"
)

func WriteJSONFile(user, host, target string, v any) error {
	path := filepath.Join(os.TempDir(), "remote", util.RandomID())
	err := util.WriteJSONFile(path, v)
	if err != nil {
		return err
	}
	cmd := exec.Command("scp", path, fmt.Sprintf("%s@%s:%s", user, host, target))
	return util.Run(cmd)
}
