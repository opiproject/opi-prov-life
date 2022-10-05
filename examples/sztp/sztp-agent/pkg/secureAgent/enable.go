package secureAgent

func RunCommandEnable() error {
	a := NewAgent("")
	err := a.execEnable()
	return err
}

func (a *Agent) execEnable() error {
	return nil
}
