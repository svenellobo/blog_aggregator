package config

import (
	"encoding/json"
	"errors"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {	
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		errors.New("Data could not be read")
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)

	return cfg, nil

}



func getConfigFilePath() (string, error) {	
	home_dir, err := os.UserHomeDir()
	if err != nil {
		errors.New("Home directory not found")
		return "", err
	}
	full_path := home_dir + "/" + configFileName
	
	return full_path, nil
}



func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	data, err := json.Marshal(c)
	if err != nil {
		errors.New("Problem with writing a file")
		return err
	}
	path, err := getConfigFilePath()
	if err != nil {
		errors.New("File not found")
		return err
	}

	os.WriteFile(path, data, 0644)

	return nil
	

}
