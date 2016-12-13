package models

type (
	// EnterpriseConfigParams struct
	EnterpriseConfigParams struct {
		DatabasePlatform     string `json:"databasePlatform"`
		ConnectString        string `json:"connectString"`
		TableOwner           string `json:"tableOwner"`
		DBUsername           string `json:"dbUsername"`
		DBUserpasswd         string `json:"dbUserpasswd"`
		ODBCDataSource       string `json:"odbcDataSource"`
		SQLDatabase          string `json:"sqlDatabase"`
		SQLServer            string `json:"sqlServer"`
		Db2DatabaseAlias     string `json:"db2DatabaseAlias"`
		Db2CurrentSQLID      string `json:"db2CurrentSQLID"`
		Encrypt              string `json:"encrypt"`
		SiebelEncryption     string `json:"siebelEncryption"`
		KeyFileName          string `json:"keyFileName"`
		KeyFilePassword      string `json:"keyFilePassword"`
		PeerAuth             string `json:"peerAuth"`
		PeerCertValidation   string `json:"peerCertValidation"`
		CACertFileName       string `json:"caCertFileName"`
		CertFileNameServer   string `json:"certFileNameServer"`
		RequestServer        string `json:"requestServer"`
		SecAdptMode          string `json:"secAdptMode"`
		ServerFileSystem     string `json:"serverFileSystem"`
		CloudRegistryAddress string `json:"cloudRegistryAddress"`
	}

	// EnterpriseProfile struct
	EnterpriseProfile struct {
		Profile                Profile                `json:"profile"`
		EnterpriseConfigParams EnterpriseConfigParams `json:"enterpriseConfigParams"`
	}
)
