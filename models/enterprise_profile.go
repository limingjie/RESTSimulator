package models

type (
	// EnterpriseConfigParams struct
	EnterpriseConfigParams struct {
		UserName           string `json:"UserName"`
		Password           string `json:"Password"`
		DBConnectString    string `json:"DBConnectString"`
		SecAdptProfileName string `json:"SecAdptProfileName"`
		PrimaryLanguage    string `json:"PrimaryLanguage"`
		DatabasePlatform   string `json:"DatabasePlatform"`
		TableOwner         string `json:"TableOwner"`
		DBUsername         string `json:"DBUsername"`
		DBUserPasswd       string `json:"DBUserPasswd"`
		SQLDatabase        string `json:"SqlDatabase"`
		SQLServer          string `json:"SqlServer"`
		SQLServerPort      string `json:"SqlServerPort"`
		Db2DatabaseAlias   string `json:"Db2DatabaseAlias"`
		Db2CurrentSQLID    string `json:"Db2CurrentSQLID"`
		Encrypt            string `json:"Encrypt"`
		KeyFileName        string `json:"KeyFileName"`
		KeyFilePassword    string `json:"KeyFilePassword"`
		PeerAuth           string `json:"PeerAuth"`
		PeerCertValidation string `json:"PeerCertValidation"`
		CACertFileName     string `json:"CACertFileName"`
		ServerFileSystem   string `json:"ServerFileSystem"`
	}

	// EnterpriseProfile struct
	EnterpriseProfile struct {
		Profile                Profile                `json:"Profile"`
		EnterpriseConfigParams EnterpriseConfigParams `json:"EnterpriseConfigParams"`
	}
)
