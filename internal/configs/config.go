package configs

import (
	"github.com/stdyum/api-common/env"
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
