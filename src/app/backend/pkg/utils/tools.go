package utils

import (
	"os"
)

// InstallCheck return install check
func InstallCheck() (bool, error) {
	cfgPath := os.Getenv("YunusCFG_PATH")
	if cfgPath != "" {
		if _, err := os.Stat(cfgPath); err != nil {
			return false, err
		}
	}
	return true, nil
}
