package config

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}


