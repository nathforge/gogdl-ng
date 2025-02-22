package env

import (
	"os"
	"path/filepath"
)

const (
	config     = "config"
	downloads  = "downloads"
	completed  = "completed"
	incomplete = "incomplete"
)

var ConfigurationFolder string
var CompletedFolder string
var IncompleteFolder string

func NewEnvironment() error {
	configurationFolder, err := getConfigurationFolder()

	if err != nil {
		return err
	}

	ConfigurationFolder = configurationFolder

	completedFolder, err := getCompletedFolder()

	if err != nil {
		return err
	}

	CompletedFolder = completedFolder

	incompleteFolder, err := getIncompleteFolder()

	if err != nil {
		return err
	}

	IncompleteFolder = incompleteFolder

	return nil
}

func getConfigurationFolder() (string, error) {
	wd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	path := filepath.Join(wd, config)

	if err := os.MkdirAll(path, 0755); err != nil {
		return "", err
	}

	return path, nil
}

func getIncompleteFolder() (string, error) {
	return getDownloadsFolderPath(incomplete)
}

func getCompletedFolder() (string, error) {
	return getDownloadsFolderPath(completed)
}

func getDownloadsFolderPath(lastPathSegment string) (string, error) {
	wd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	path := filepath.Join(wd, downloads, lastPathSegment)

	if err := os.MkdirAll(path, 0755); err != nil {
		return "", err
	}

	return path, nil
}
