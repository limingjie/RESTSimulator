package models

import (
	"../logger"
)

type (
	// Deployment struct
	Deployment struct {
		PhysicalHostIP string `json:"physicalHostIP"`
		ProfileName    string `json:"profileName"`
		Action         string `json:"action"`
		Status         string `json:"status"`
	}
)

// Check Deployment
func (deployment *Deployment) Check() {
	if deployment.Status == "" {
		deployment.Status = "Deployed"
	}

	logger.Logger("Deployment.Check", "Deployed")
}
