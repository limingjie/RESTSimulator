package models

import (
	"encoding/json"

	"../logger"
)

type (
	// Profile struct
	Profile struct {
		ProfileName      string `json:"ProfileName"`
		AccessPermission string `json:"AccessPermission"`
		LastUpdated      string `json:"LastUpdated"`
		count            int
	}

	// LogProperties struct
	LogProperties struct {
		LogLevel string `json:"LogLevel"`
	}

	// GatewayIdentity struct
	GatewayIdentity struct {
		GatewayHost string `json:"GatewayHost"`
		GatewayPort string `json:"GatewayPort"`
		AuthToken   string `json:"AuthToken"`
	}
)

// Deploy - Deploy the profile
func (profile *Profile) Deploy() {
	profile.count++
	profile.AccessPermission = "ReadOnly"

	msg, _ := json.Marshal(profile)
	logger.Logger("deployProfile", string(msg))
}

// Undeploy - Undeploy the profile
func (profile *Profile) Undeploy() {
	profile.count--
	if profile.count == 0 {
		profile.AccessPermission = "ReadWrite"
	}

	msg, _ := json.Marshal(profile)
	logger.Logger("undeployProfile", string(msg))
}
