package env

func Fill(s any) error {
	err = assignStructFromMap("", s, envs)
	return err
}
