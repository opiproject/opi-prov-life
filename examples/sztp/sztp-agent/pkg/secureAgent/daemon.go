package secureAgent

func RunCommandDaemon() error {
	a := NewAgent("")
	err := a.execDaemon()
	return err
}

func (a *Agent) execDaemon() error {
	return nil
}
