package models

type (
	// SWSMDeployParams struct
	SWSMDeployParams struct {
		SWSMNode     string `json:"Node"`
		SWSMNodeDesc string `json:"NodeDesc"`
	}

	// SWSMDeployment struct
	SWSMDeployment struct {
		Deployment       Deployment       `json:"DeploymentInfo"`
		SWSMDeployParams SWSMDeployParams `json:"DeploymentParam"`
	}
)
