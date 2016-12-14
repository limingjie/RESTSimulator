package models

type (
	// CacheServerDeployParams struct
	CacheServerDeployParams struct {
		CacheServerAgentNode     string `json:"cacheServerAgentNode"`
		CacheServerAgentNodeDesc string `json:"cacheServerAgentNodeDesc"`
	}

	// CacheServerDeployment struct
	CacheServerDeployment struct {
		Deployment              Deployment              `json:"deploymentInfo"`
		CacheServerDeployParams CacheServerDeployParams `json:"cacheServerDeployParams"`
	}
)
