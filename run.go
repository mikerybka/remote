package remote

import "fmt"

func Run(user, host, cmd string) error {
	c := &Command{
		Cmd: cmd,
	}
	out, err := c.Run(&Machine{
		UserID: user,
		Host:   host,
	})
	if err != nil {
		fmt.Println(string(out))
	}
	return err
}
