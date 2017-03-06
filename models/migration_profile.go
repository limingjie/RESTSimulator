package models

type (
	// MigrationConfigParams struct
	MigrationConfigParams struct {
		DatabaseType          string `json:"DatabaseType"`
		Hostname              string `json:"Hostname"`
		Portnum               string `json:"Portnum"`
		TableOwner            string `json:"TableOwner"`
		Username              string `json:"Username"`
		Password              string `json:"Password"`
		ServiceName           string `json:"ServiceName"`
		Instance              string `json:"Instance"`
		DatabaseName          string `json:"DatabaseName"`
		AuthenticationType    string `json:"AuthenticationType"`
		AuthenticationHost    string `json:"AuthenticationHost"`
		UserSpec              string `json:"UserSpec"`
		AssertionSpec         string `json:"AssertionSpec"`
		Timeout               int    `json:"Timeout"`
		SleepTime             int    `json:"SleepTime"`
		LogLevel              string `json:"LogLevel"`
		SiebelApplicationName string `json:"SiebelApplicationName"`
		Language              string `json:"Language"`
	}

	// MigrationProfile struct
	MigrationProfile struct {
		Profile               Profile               `json:"Profile"`
		MigrationConfigParams MigrationConfigParams `json:"MigrationConfigParams"`
	}
)
