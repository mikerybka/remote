package remote

import (
	"fmt"
	"os/exec"

	"github.com/mikerybka/util"
)

func Copy(user, host, dst, src string) error {
	cmd := exec.Command("scp", src, fmt.Sprintf("%s@%s:%s", user, host, src))
	return util.Run(cmd)
}
