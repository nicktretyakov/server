package testutil

import (
	"os"

	"gopkg.in/yaml.v3"
)

func readFile(fileName string, dest interface{}) error {
	f, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(f, dest)
}
