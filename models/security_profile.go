package models

type (
	// DataSource struct
	DataSource struct {
		Name                        string `json:"Name"`
		Type                        string `json:"Type"`
		Host                        string `json:"Host"`
		Port                        int    `json:"Port"`
		SQLStyle                    string `json:"SqlStyle"`
		Endpoint                    string `json:"Endpoint"`
		TableOwner                  string `json:"TableOwner"`
		BaseDN                      string `json:"BaseDN"`
		SharedCredentialsDN         string `json:"SharedCredentialsDN"`
		SharedDBPassword            string `json:"SharedDBPassword"`
		SharedDBUsername            string `json:"SharedDBUsername"`
		UsernameAttributeType       string `json:"UsernameAttributeType"`
		PasswordAttributeType       string `json:"PasswordAttributeType"`
		CredentialsAttributeType    string `json:"CredentialsAttributeType"`
		RolesAttributeType          string `json:"RolesAttributeType"`
		SaltAttributeType           string `json:"SaltAttributeType"`
		ApplicationUser             string `json:"ApplicationUser"`
		ApplicationPassword         string `json:"ApplicationPassword"`
		HashDBPwd                   bool   `json:"HashDBPwd"`
		HashUserPwd                 bool   `json:"HashUserPwd"`
		SaltUserPwd                 bool   `json:"SaltUserPwd"`
		HashAlgorithm               string `json:"HashAlgorithm"`
		PropagateChange             bool   `json:"PropagateChange"`
		CRC                         bool   `json:"CRC"`
		SingleSignOn                bool   `json:"SingleSignOn"`
		TrustToken                  string `json:"TrustToken"`
		UseAdapterUsername          bool   `json:"UseAdapterUsername"`
		SiebelUsernameAttributeType string `json:"SiebelUsernameAttributeType"`
		SSL                         bool   `json:"SSL"`
		WalletPassword              string `json:"WalletPassword"`
	}

	// SecurityConfigParams struct
	SecurityConfigParams struct {
		DataSources                      []DataSource `json:"DataSources"`
		NSAdminRole                      []string     `json:"NSAdminRole"`
		DBSecurityAdapterDataSource      string       `json:"DBSecurityAdapterDataSource"`
		DBSecurityAdapterPropagateChange bool         `json:"DBSecurityAdapterPropagateChange"`
		SecAdptName                      string       `json:"SecAdptName"`
		SecAdptMode                      string       `json:"SecAdptMode"`
		TestUserName                     string       `json:"TestUserName"`
		TestUserPwd                      string       `json:"TestUserPwd"`
	}

	// SecurityProfile struct
	SecurityProfile struct {
		Profile              Profile              `json:"Profile"`
		SecurityConfigParams SecurityConfigParams `json:"SecurityConfigParams"`
	}
)
