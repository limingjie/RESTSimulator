package models

type (
	// ConstraintEngineConfigParams struct
	ConstraintEngineConfigParams struct {
		CGHostURL          string `json:"CGHostURL"`
		CacheClientProfile string `json:"CacheClientProfile"`
		LogLocation        string `json:"LogLocation"`
		LogFileName        string `json:"LogFileName"`
		DefaultLogLevel    string `json:"DefaultLogLevel"`
	}

	// ConstraintEngineProfile struct
	ConstraintEngineProfile struct {
		Profile                      Profile                      `json:"Profile"`
		ConstraintEngineConfigParams ConstraintEngineConfigParams `json:"ConstraintEngineConfigParams"`
	}
)
