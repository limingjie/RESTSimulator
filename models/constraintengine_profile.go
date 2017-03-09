package models

type (
	// ConstraintEngineConfigParams struct
	ConstraintEngineConfigParams struct {
		CacheClientProfile string `json:"CacheClientProfile"`
		SecurityProfile    string `json:"SecurityProfile"`
		DefaultLogLevel    string `json:"DefaultLogLevel"`
	}

	// ConstraintEngineProfile struct
	ConstraintEngineProfile struct {
		Profile                      Profile                      `json:"Profile"`
		ConstraintEngineConfigParams ConstraintEngineConfigParams `json:"ConstraintEngineConfigParams"`
	}
)
