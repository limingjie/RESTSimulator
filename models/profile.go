package models

type (
	// Profile struct
	Profile struct {
		ProfileName      string `json:"profileName"`
		AccessPermission string `json:"accessPermission"`
		LastUpdated      string `json:"lastUpdated"`
	}
)
