package internal

func App() error {
	handler, _, err := Configure()
	if err != nil {
		return err
	}

	_ = handler
	return nil
}
