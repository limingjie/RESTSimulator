package models

type (
	// ServerConfigParams struct
	ServerConfigParams struct {
		SiebelEnterprise           string `json:"SiebelEnterprise"`
		ModifyServerEncrypt        bool   `json:"ModifyServerEncrypt"`
		ModifyServerAuth           bool   `json:"ModifyServerAuth"`
		ClusteringEnvironmentSetup string `json:"ClusteringEnvironmentSetup"`
		SiebelClusterGateway       string `json:"SiebelClusterGateway"`
		UseOracleConnector         string `json:"UseOracleConnector"`
		Encrypt                    string `json:"Encrypt"`
		CACertFileName             string `json:"CACertFileName"`
		CertFileNameServer         string `json:"CertFileNameServer"`
		Username                   string `json:"UserName"`
		Password                   string `json:"Password"`
		NameserverHostName         string `json:"NameserverHostName"`
		NamesrvrPort               string `json:"NamesrvrPort"`
		SQLServerPort              string `json:"SqlServerPort"`
	}

	// ServerProfile struct
	ServerProfile struct {
		Profile            Profile            `json:"Profile"`
		ServerConfigParams ServerConfigParams `json:"ServerConfigParams"`
	}
)
