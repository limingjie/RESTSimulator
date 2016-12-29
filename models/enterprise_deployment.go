package models

type (
	// EnterpriseDeployParams struct
	EnterpriseDeployParams struct {
		SiebelEnterprise string `json:"SiebelEnterprise"`
		EnterpriseDesc   string `json:"EnterpriseDesc"`
	}

	// EnterpriseDeployment struct
	EnterpriseDeployment struct {
		Deployment             Deployment             `json:"DeploymentInfo"`
		EnterpriseDeployParams EnterpriseDeployParams `json:"EnterpriseDeployParams"`
	}
)
