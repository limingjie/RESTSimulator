package models

type (
	// GatewayClusterNodeRegistryList struct
	GatewayClusterNodeRegistryList struct {
		PhysicalHostIP string `json:"PhysicalHostIP"`
		RegistryStatus string `json:"RegistryStatus"`
	}

	// GatewayClusterNodeServiceList struct
	GatewayClusterNodeServiceList struct {
		PhysicalHostIP string `json:"PhysicalHostIP"`
		ServiceStatus  string `json:"ServiceStatus"`
	}

	// GatewayClusterDeployment struct
	GatewayClusterDeployment struct {
		Deployment                     Deployment                       `json:"DeploymentInfo"`
		GatewayClusterNodeRegistryList []GatewayClusterNodeRegistryList `json:"GatewayClusterNodeRegistryList"`
		GatewayClusterNodeServiceList  []GatewayClusterNodeServiceList  `json:"GatewayClusterNodeServiceList"`
	}
)
