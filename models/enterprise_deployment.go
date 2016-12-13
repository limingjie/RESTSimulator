package models

type (
	// EnterpriseDeployParams struct
	EnterpriseDeployParams struct {
		EnterpriseServer string `json:"enterpriseServer"`
		EnterpriseDesc   string `json:"enterpriseDesc"`
	}

	// EnterpriseDeployment struct
	EnterpriseDeployment struct {
		Deployment             Deployment             `json:"deploymentInfo"`
		EnterpriseDeployParams EnterpriseDeployParams `json:"enterpriseDeployParams"`
	}
)
