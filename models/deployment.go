package models

type (
	// Deployment struct
	Deployment struct {
		PhysicalHostIP string `json:"physicalHostIP"`
		ProfileName    string `json:"profileName"`
		Action         string `json:"action"`
		Status         string `json:"status"`
	}
)
