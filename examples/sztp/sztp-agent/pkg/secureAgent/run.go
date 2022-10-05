package secureAgent

func RunCommandRun() error {
	a := NewAgent("")
	err := a.run()
	return err
}

func (a *Agent) run() error {
	return nil
}
