package config

import (
	"encoding/json"
	"os"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + "\\" + configFileName, nil
}

func (cfg Config) write() error {
	filePath, err := getConfigFilePath()
	if err != nil { return err }


	jsonData, err := json.Marshal(cfg)
	if err != nil { return err }

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil { return err }

	return nil
}
