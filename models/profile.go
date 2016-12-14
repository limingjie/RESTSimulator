package models

import (
	"encoding/json"

	"../logger"
)

type (
	// Profile struct
	Profile struct {
		ProfileName      string `json:"profileName"`
		AccessPermission string `json:"accessPermission"`
		LastUpdated      string `json:"lastUpdated"`
		Count            int    `json:"count"`
	}
)

// Deploy - Deploy the profile
func (profile *Profile) Deploy() {
	profile.Count++
	profile.AccessPermission = "READONLY"

	msg, _ := json.Marshal(profile)
	logger.Logger("deployProfile", string(msg))
}

// Undeploy - Undeploy the profile
func (profile *Profile) Undeploy() {
	profile.Count--
	if profile.Count == 0 {
		profile.AccessPermission = "READWRITE"
	}

	msg, _ := json.Marshal(profile)
	logger.Logger("undeployProfile", string(msg))
}
