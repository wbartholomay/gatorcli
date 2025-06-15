package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error){
	filePath, err := getConfigFilePath()
	if err != nil { return Config{},err }

	file, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("config file does not exist at %s: %w", filePath, err)
	}

	var jsonData Config

	err = json.Unmarshal(file, &jsonData)
	if err != nil {
		return Config{},fmt.Errorf("error unmarshaling config file: %w", err)
	}

	return jsonData, nil
}

func (c Config) SetUser(username string) error {
	c.CurrentUsername = username
	err := c.write()
	if err != nil { return err }

	return nil
}