package load

import (
	"fmt"
	"os"
)

func ReadConfig(configPath string) error {
	_, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	return nil
}
