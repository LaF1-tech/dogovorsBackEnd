package env

import (
	"os"
	"strings"
)

var err error
var envs map[string]string

func parseEnvData(content string) map[string]string {
	lines := strings.Split(content, "\n")
	envs := make(map[string]string)
	for _, line := range lines {
		line = strings.ReplaceAll(strings.TrimSpace(line), "\r", "")
		i := strings.Index(line, "=")
		if i == -1 {
			continue
		}

		name := line[:i]
		value := line[i+1:]
		if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
			value = value[1 : len(value)-1]
		}
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
			value = value[1 : len(value)-1]
		}

		envs[name] = value
	}

	return envs
}

func parseEnvs() {
	data := strings.Join(os.Environ(), "\n")
	envs = parseEnvData(data)
}

func init() {
	err = readEnvFile()
	parseEnvs()
}
