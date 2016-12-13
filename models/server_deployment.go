package models

type (
	// ServerDeployParams struct
	ServerDeployParams struct {
		SiebelServer        string `json:"siebelServer"`
		SiebelServerDesc    string `json:"siebelServerDesc"`
		SelectOMSIA         string `json:"selectOMSIA"`
		EnableCompGroupsSIA string `json:"enableCompGroupsSIA"`
		SCBPort             string `json:"sCBPort"`
		LocalSynchMgrPort   string `json:"localSynchMgrPort"`
	}

	// ServerDeployment struct
	ServerDeployment struct {
		Deployment         Deployment         `json:"deploymentInfo"`
		ServerDeployParams ServerDeployParams `json:"enterpriseDeployParams"`
	}
)
