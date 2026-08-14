package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type DefaultConfig struct {
	Theme  string `json:"theme"`
	Volume int    `json:"volume"`
}

type Config struct {
	Theme  string `json:"theme"`
	Volume int    `json:"volume"`
}

var Defaults = DefaultConfig{
	Theme:  "default",
	Volume: 100,
}

func getConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(configDir, "ctread")

	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(appDir, "config.json"), nil
}

func Load() (Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configPath)
	if os.IsNotExist(err) {
		config := Config{
			Theme:  Defaults.Theme,
			Volume: Defaults.Volume,
		}

		err = Save(config)
		if err != nil {
			return Config{}, err
		}

		return config, nil
	}

	if err != nil {
		return Config{}, err
	}

	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func Save(config Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func Reset() error {
	config := Config{
		Theme:  Defaults.Theme,
		Volume: Defaults.Volume,
	}

	return Save(config)
}

func Setting(setting string, value ...any) error {
	config, err := Load()
	if err != nil {
		return err
	}

	if len(value) == 0 {
		switch setting {
		case "theme":
			fmt.Println(config.Theme)

		case "volume":
			fmt.Println(config.Volume)

		default:
			return fmt.Errorf("unknown setting: %s", setting)
		}

		return nil
	}

	switch setting {
	case "theme":
		config.Theme = fmt.Sprint(value[0])

	case "volume":
		var volume int

		switch v := value[0].(type) {
		case int:
			volume = v

		case string:
			_, err := fmt.Sscanf(v, "%d", &volume)
			if err != nil {
				return fmt.Errorf("volume must be a number")
			}

		default:
			return fmt.Errorf("volume must be a number")
		}

		config.Volume = volume

	default:
		return fmt.Errorf("unknown setting: %s", setting)
	}

	err = Save(config)
	if err != nil {
		return err
	}

	fmt.Printf("%s updated\n", setting)

	return nil
}
