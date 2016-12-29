package models

type (
	// CacheServerDeployParams struct
	CacheServerDeployParams struct {
		CacheServerAgentNode     string `json:"CacheServerAgentNode"`
		CacheServerAgentNodeDesc string `json:"CacheServerAgentNodeDesc"`
	}

	// CacheServerDeployment struct
	CacheServerDeployment struct {
		Deployment              Deployment              `json:"DeploymentInfo"`
		CacheServerDeployParams CacheServerDeployParams `json:"CacheServerDeployParams"`
	}
)
