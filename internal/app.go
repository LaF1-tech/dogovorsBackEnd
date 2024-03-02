package internal

func App() error {
	handler, _, err := Configure()
	if err != nil {
		return err
	}

	return handler.Routes().Run()
}
