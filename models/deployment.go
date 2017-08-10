package models

import (
	"log"
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
	switch deployment.Action {
	case "Save":
		deployment.Status = "Saved"
	case "Deployed":
		fallthrough
	default:
		deployment.Status = "Deployed"
	}

	log.Println("Deployment.Check", deployment.Status)
}
