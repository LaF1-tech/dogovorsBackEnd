package configs

import (
	"dogovorsBackEnd/pkg/env"
)

var Config Model

type Model struct {
	Database DatabaseConfig `env:"DATABASE"`
}

func init() {
	err := env.Fill(&Config)
	if err != nil {
		panic("cannot fill config: " + err.Error())
	}
}
