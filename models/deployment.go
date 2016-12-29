package models

import (
	"../logger"
)

type (
	// Deployment struct
	Deployment struct {
		PhysicalHostIP string `json:"PhysicalHostIP"`
		ProfileName    string `json:"ProfileName"`
		Action         string `json:"Action"`
		Status         string `json:"Status"`
	}
)

// Check Deployment
func (deployment *Deployment) Check() {
	if deployment.Status == "" {
		deployment.Status = "Deployed"
	}

	logger.Logger("Deployment.Check", "Deployed")
}
