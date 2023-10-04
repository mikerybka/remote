package remote

func NewMachine(userID, host string) *Machine {
	return &Machine{
		UserID: userID,
		Host:   host,
	}
}

type Machine struct {
	UserID string
	Host   string
}
