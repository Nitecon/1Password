package user

import (
	"fmt"
	"os"
	"os/user"
)

func GetConfPath() (string, error) {
	cu, err := user.Current()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s%s", cu.HomeDir, string(os.PathSeparator), ".1Password-config.json"), nil
}
