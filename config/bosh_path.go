package config

import (
	"fmt"
	"os"
	"os/exec"
)

func resolveBoshFromEnvironment() string {
	boshExecutable := "bosh"
	if len(os.Getenv("BBL_BOSH_BIN")) > 0 {
		boshExecutable = os.Getenv("BBL_BOSH_BIN")
	}
	return boshExecutable
}

func GetBOSHPath() (string, error) {
	var boshPath = resolveBoshFromEnvironment()

	path, err := exec.LookPath("bosh2")
	if err != nil {
		if err.(*exec.Error).Err != exec.ErrNotFound {
			return "", fmt.Errorf("failed when searching for BOSH: %s", err) // not tested
		}
	}

	if path != "" {
		boshPath = path
	}

	return boshPath, nil
}
