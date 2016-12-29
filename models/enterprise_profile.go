package models

type (
	// EnterpriseConfigParams struct
	EnterpriseConfigParams struct {
		UserName                       string `json:"UserName"`
		Password                       string `json:"Password"`
		NameserverHostName             string `json:"NameserverHostName"`
		NamesrvrPort                   string `json:"NamesrvrPort"`
		DBConnectString                string `json:"DBConnectString"`
		SecAdptName                    string `json:"SecAdptName"`
		UseAdapterUserName             bool   `json:"UseAdapterUserName"`
		PropagateGatewayAuthentication bool   `json:"PropagateGatewayAuthentication"`
		ChartImageFormat               string `json:"ChartImageFormat"`
		ChartServerHost                string `json:"ChartServerHost"`
		BaseDN                         string `json:"BaseDN"`
		CRC                            string `json:"CRC"`
		FQDNName                       string `json:"FQDNName"`
		CredentialsAttributeType       string `json:"CredentialsAttributeType"`
		Port                           string `json:"Port"`
		PortMode                       string `json:"PortMode"`
		PropagateChange                bool   `json:"PropagateChange"`
		RolesAttributeType             string `json:"RolesAttributeType"`
		UsernameAttributeType          string `json:"UsernameAttributeType"`
		SingleSignOn                   bool   `json:"SingleSignOn"`
		SharedCredentialsDN            string `json:"SharedCredentialsDN"`
		ApplicationUser                string `json:"ApplicationUser"`
		ApplicationPassword            string `json:"ApplicationPassword"`
		DatabasePlatform               string `json:"DatabasePlatform"`
		TableOwner                     string `json:"TableOwner"`
		DBUsername                     string `json:"DBUsername"`
		DBUserPasswd                   string `json:"DBUserPasswd"`
		SQLDatabase                    string `json:"SqlDatabase"`
		SQLServer                      string `json:"SqlServer"`
		SQLServerPort                  string `json:"SqlServerPort"`
		Db2DatabaseAlias               string `json:"Db2DatabaseAlias"`
		Db2CurrentSQLID                string `json:"Db2CurrentSQLID"`
		Encrypt                        string `json:"Encrypt"`
		KeyFileName                    string `json:"KeyFileName"`
		KeyFilePassword                string `json:"KeyFilePassword"`
		PeerAuth                       string `json:"PeerAuth"`
		PeerCertValidation             string `json:"PeerCertValidation"`
		CACertFileName                 string `json:"CACertFileName"`
		SecAdptMode                    string `json:"SecAdptMode"`
		ServerFileSystem               string `json:"ServerFileSystem"`
	}

	// EnterpriseProfile struct
	EnterpriseProfile struct {
		Profile                Profile                `json:"Profile"`
		EnterpriseConfigParams EnterpriseConfigParams `json:"EnterpriseConfigParams"`
	}
)
