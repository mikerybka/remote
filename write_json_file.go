package remote

import (
	"os"
	"path/filepath"

	"github.com/mikerybka/util"
)

func WriteJSONFile(user, host, dst string, v any) error {
	src := filepath.Join(os.TempDir(), "remote", util.RandomID())
	err := util.WriteJSONFile(src, v)
	if err != nil {
		return err
	}
	return Copy(user, host, dst, src)
}
