package config

import (
	"encoding/base64"
	"encoding/json"
	"os"
)

// EnvVarsStruct contains self defined env vars
type EnvVarsStruct struct {
	GeniusAccessToken string `json:"GENIUS_ACCESS_TOKEN"`
}

// EnvVars contains parsed env vars
var EnvVars EnvVarsStruct

// InitEnvVars reads env vars in EnvVars
func InitEnvVars() error {
	platformVars := os.Getenv("PLATFORM_VARIABLES")
	jsonVars, err := base64.StdEncoding.DecodeString(platformVars)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(jsonVars, &EnvVars); err != nil {
		return err
	}
	return nil
}
