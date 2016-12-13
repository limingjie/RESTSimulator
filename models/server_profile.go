package models

type (
	// ServerConfigParams struct
	ServerConfigParams struct {
		SiebelEnterprise           string `json:"siebelEnterprise"`
		AnonLoginUserName          string `json:"anonLoginUserName"`
		AnonLoginPassword          string `json:"anonLoginPassword"`
		ModifyServerEncrypt        string `json:"modifyServerEncrypt"`
		ModifyServerAuth           string `json:"modifyServerAuth"`
		ClusteringEnvironmentSetup string `json:"clusteringEnvironmentSetup"`
		SiebelClusterGateway       string `json:"siebelClusterGateway"`
		UseOracleConnector         string `json:"useOracleConnector"`
		Encrypt                    string `json:"encrypt"`
		CACertFileName             string `json:"cACertFileName"`
		CertFileNameServer         string `json:"certFileNameServer"`
		UserName                   string `json:"userName"`
		Password                   string `json:"password"`
		NameserverHostName         string `json:"nameserverHostName"`
		NamesrvrPort               string `json:"namesrvrPort"`
	}

	// ServerProfile struct
	ServerProfile struct {
		Profile            Profile            `json:"profile"`
		ServerConfigParams ServerConfigParams `json:"serverConfigParams"`
	}
)
