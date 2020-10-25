package agent

type rsyslog struct {
	conf map[string]string
}

func (r *rsyslog)Start()bool{
	return true
}

func (r *rsyslog)Stop()bool{
	return true
}

func (r *rsyslog)Restart()bool{
	return true
}
