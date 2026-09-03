package importers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	Styles "github.com/Hollowzzzy/CTread/lipgloss"
)

var registry []BookFormat

type BookFormat struct {
	Name     string
	Path     string
	Modified time.Time
}

func FindBook(name string) (*BookFormat, bool) {
	_, err := Load()
	if err != nil {
		return nil, false
	}

	for i := range registry {
		if registry[i].Name == name {
			return &registry[i], true
		}
	}

	return nil, false
}

func RegistryAdd(name string, path string, modified time.Time) error {
	_, er := Load()
	if er != nil {
		return er
	}
	book := BookFormat{
		Name:     name,
		Path:     path,
		Modified: modified,
	}

	registry = append(registry, book)

	err := Save(registry)
	if err != nil {
		return err
	}

	return nil
}

func RegistryDelete(name string) error {
	Load()
	for i, book := range registry {
		if book.Name == name {
			registry = append(registry[:i], registry[i+1:]...)

			if err := Save(registry); err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("book %q was not found in the registry", name)
}

func RegistryClear() error {
	registry = registry[:0]
	fmt.Println(Styles.INFO.Render("The registry has been cleared!"))
	return Save(registry)
}

func getRegistryPath() (string, error) {
	registryDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(registryDir, "ctread")

	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(appDir, "registry.json"), nil
}

func Save(registry []BookFormat) error {
	registryPath, err := getRegistryPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(registry, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(registryPath, data, 0644)
}

func Load() ([]BookFormat, error) {
	registryPath, err := getRegistryPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(registryPath)
	if os.IsNotExist(err) {
		registry = []BookFormat{}

		err = Save(registry)
		if err != nil {
			return nil, err
		}

		return registry, nil
	}

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &registry)
	if err != nil {
		return nil, err
	}

	return registry, nil
}
