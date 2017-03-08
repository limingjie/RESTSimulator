package models

type (
	// ConstraintEngineDeployParams struct
	ConstraintEngineDeployParams struct {
		DeploymentName string `json:"DeploymentName"`
	}

	// ConstraintEngineDeployment struct
	ConstraintEngineDeployment struct {
		Deployment                   Deployment                   `json:"DeploymentInfo"`
		ConstraintEngineDeployParams ConstraintEngineDeployParams `json:"ConstraintEngineDeployParams"`
	}
)
