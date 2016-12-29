package models

type (
	// ServerDeployParams struct
	ServerDeployParams struct {
		AnonLoginUserName   string `json:"AnonLoginUserName"`
		AnonLoginPassword   string `json:"AnonLoginPassword"`
		SiebelServer        string `json:"SiebelServer"`
		SiebelServerDesc    string `json:"SiebelServerDesc"`
		SelectOMSIA         string `json:"SelectOMSIA"`
		EnableCompGroupsSIA string `json:"EnableCompGroupsSIA"`
		SCBPort             string `json:"SCBPort"`
		LocalSynchMgrPort   string `json:"LocalSynchMgrPort"`
	}

	// ServerDeployment struct
	ServerDeployment struct {
		Deployment         Deployment         `json:"DeploymentInfo"`
		ServerDeployParams ServerDeployParams `json:"ServerDeployParams"`
	}
)
