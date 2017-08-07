package models

type (
	// GatewayClusterConfigParams struct
	GatewayClusterConfigParams struct {
		ClientPort   string `json:"ClientPort"`
		FollowerPort string `json:"FollowerPort"`
		LeaderPort   string `json:"LeaderPort"`
	}

	// GatewayClusterProfile struct
	GatewayClusterProfile struct {
		Profile                    Profile                    `json:"Profile"`
		GatewayClusterConfigParams GatewayClusterConfigParams `json:"GatewayClusterConfigParams"`
	}
)
