package models

type (
	// GatewayClusterNodeRegistry struct
	GatewayClusterNodeRegistry struct {
		PhysicalHostIP string `json:"PhysicalHostIP"`
		RegistryStatus string `json:"RegistryStatus"`
	}

	// GatewayClusterNodeService struct
	GatewayClusterNodeService struct {
		PhysicalHostIP string `json:"PhysicalHostIP"`
		ServiceStatus  string `json:"ServiceStatus"`
	}

	// GatewayClusterDeployment struct
	GatewayClusterDeployment struct {
		Deployment                 Deployment                   `json:"DeploymentInfo"`
		GatewayClusterNodeRegistry []GatewayClusterNodeRegistry `json:"GatewayClusterNodeRegistry"`
		GatewayClusterNodeService  []GatewayClusterNodeService  `json:"GatewayClusterNodeService"`
	}
)
