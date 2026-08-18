package functionality

import (
	"ctread/config"
	"fmt"
)

func ConfigFunc(args []string) {
	var err error
	if args[0] != "reset" {
		if len(args) == 1 {
			err = config.Setting(args[0])
		} else {
			err = config.Setting(args[0], args[1])
		}
	} else {
		err = config.Reset()
	}
	if err != nil {
		fmt.Println("Error:", err)
	}
}
