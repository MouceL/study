package agent

type Controller interface {
	Start() bool
	Stop() bool
	Restart() bool
}



