package models

type (
	// ServerDeployParams struct
	ServerDeployParams struct {
		SiebelServer     string `json:"SiebelServer"`
		SiebelServerDesc string `json:"SiebelServerDesc"`
		DeployedLanguage string `json:"DeployedLanguage"`
	}

	// ServerDeployment struct
	ServerDeployment struct {
		Deployment         Deployment         `json:"DeploymentInfo"`
		ServerDeployParams ServerDeployParams `json:"ServerDeployParams"`
	}
)
