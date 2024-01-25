package handlers

import (
	"encoding/json"
	"io"
	"os"
)

type ConfigHandler struct {
	InstallerVersion string            `json:"installer-version"`
	GameDirectories  map[string]string `json:"game-directories"`
}

func MustInitConfigHandler() ConfigHandler {
	config, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	defer config.Close()

	var configHandler ConfigHandler

	readConfig, err := io.ReadAll(config)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(readConfig, &configHandler)

	if err != nil {
		panic(err)
	}

	return configHandler
}
