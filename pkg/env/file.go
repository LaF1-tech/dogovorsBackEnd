package env

import (
	"os"
)

func readEnvFile() error {
	bytes, err := os.ReadFile(".env")
	if err != nil {
		return err
	}

	content := string(bytes)
	envs := parseEnvData(content)
	for key, value := range envs {
		_ = os.Setenv(key, value)
	}

	return nil
}
