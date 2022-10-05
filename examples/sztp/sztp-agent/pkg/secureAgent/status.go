package secureAgent

func RunCommandStatus() error {
	a := NewAgent("")
	err := a.execStatus()
	return err
}

func (a *Agent) execStatus() error {
	return nil
}
