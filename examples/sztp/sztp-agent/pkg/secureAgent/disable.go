package secureAgent

func RunCommandDisable() error {
	a := NewAgent("")
	err := a.execDisable()
	return err
}

func (a *Agent) execDisable() error {
	return nil
}
