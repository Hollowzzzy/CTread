package config

import "fmt"

func UpdateSetting(setting string, value ...any) {
	if value != nil {
		fmt.Print(setting, value)
	} else {
		fmt.Print(setting)
	}
}
