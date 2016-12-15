package models

type (
	// SWSMDeployParams struct
	SWSMDeployParams struct {
		SWSMNode     string `json:"swsmNode"`
		SWSMNodeDesc string `json:"swsmNodeDesc"`
	}

	// SWSMDeployment struct
	SWSMDeployment struct {
		Deployment       Deployment       `json:"deploymentInfo"`
		SWSMDeployParams SWSMDeployParams `json:"swsmDeployParams"`
	}
)
