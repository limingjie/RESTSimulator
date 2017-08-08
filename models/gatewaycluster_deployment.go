package models

type (
	// GatewayClusterNodeRegistryList struct
	GatewayClusterNodeRegistryList struct {
		PhysicalHostIP string `json:"PhysicalHostIP"`
		Status string `json:"Status"`
	}

	// GatewayClusterNodeServiceList struct
	GatewayClusterNodeServiceList struct {
		PhysicalHostIP string `json:"PhysicalHostIP"`
		Status  string `json:"Status"`
	}

	// GatewayClusterDeployment struct
	GatewayClusterDeployment struct {
		Deployment                     Deployment                       `json:"DeploymentInfo"`
		GatewayClusterNodeRegistryList []GatewayClusterNodeRegistryList `json:"GatewayClusterNodeRegistryList"`
		GatewayClusterNodeServiceList  []GatewayClusterNodeServiceList  `json:"GatewayClusterNodeServiceList"`
	}
)
