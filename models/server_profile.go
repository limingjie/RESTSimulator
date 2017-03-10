package models

type (
	// ServerConfigParams struct
	ServerConfigParams struct {
		AnonLoginUserName          string `json:"AnonLoginUserName"`
		AnonLoginPassword          string `json:"AnonLoginPassword"`
		EnableCompGroupsSIA        string `json:"EnableCompGroupsSIA"`
		SCBPort                    string `json:"SCBPort"`
		LocalSynchMgrPort          string `json:"LocalSynchMgrPort"`
		ModifyServerEncrypt        bool   `json:"ModifyServerEncrypt"`
		ModifyServerAuth           bool   `json:"ModifyServerAuth"`
		ClusteringEnvironmentSetup string `json:"ClusteringEnvironmentSetup"`
		SiebelClusterGateway       string `json:"SiebelClusterGateway"`
		UseOracleConnector         bool   `json:"UseOracleConnector"`
		Encrypt                    string `json:"Encrypt"`
		CACertFileName             string `json:"CACertFileName"`
		CertFileNameServer         string `json:"CertFileNameServer"`
		Username                   string `json:"Username"`
		Password                   string `json:"Password"`
		SQLServerPort              string `json:"SqlServerPort"`
	}

	// ServerProfile struct
	ServerProfile struct {
		Profile            Profile            `json:"Profile"`
		ServerConfigParams ServerConfigParams `json:"ServerConfigParams"`
	}
)
