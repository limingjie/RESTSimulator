package models

type (
	// ServerConfigParams struct
	ServerConfigParams struct {
		SiebelServer     string `json:"SiebelServer"`
		SiebelServerDesc string `json:"SiebelServerDesc"`
		DeployedLanguage string `json:"DeployedLanguage"`
	}

	// ServerProfile struct
	ServerProfile struct {
		Profile            Profile            `json:"Profile"`
		ServerConfigParams ServerConfigParams `json:"ServerConfigParams"`
	}
)
