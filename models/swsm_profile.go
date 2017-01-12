package models

/* SWSM 1.0
type (
	// SWE struct
	SWE struct {
		Language             string `json:"Language"`
		SessionMonitor       string `json:"SessionMonitor"`
		AllowStats           string `json:"AllowStats"`
		ClientRootDir        string `json:"ClientRootDir"`
		MaxQueryStringLength string `json:"MaxQueryStringLength"`
		SeedFile             string `json:"SeedFile"`
	}

	// ConnMgmt struct
	ConnMgmt struct {
		EnableVirtualHosts      string `json:"EnableVirtualHosts"`
		VirtualHostsFileContent string `json:"VirtualHostsFileContent"`
		CertFileName            string `json:"CertFileName"`
		CACertFileName          string `json:"CACertFileName"`
		KeyFileName             string `json:"KeyFileName"`
		KeyFilePassword         string `json:"KeyFilePassword"`
		PeerAuth                string `json:"PeerAuth"`
		PeerCertValidation      string `json:"PeerCertValidation"`
	}

	// AuthenticationProperties struct
	AuthenticationProperties struct {
		EncryptedPassword         bool   `json:"EncryptedPassword"`
		EncryptSessionID          bool   `json:"EncryptSessionId"`
		AnonymousUserName         string `json:"AnonymousUserName"`
		AnonymousPassword         string `json:"AnonymousPassword"`
		ClientCertificate         bool   `json:"ClientCertificate"`
		TrustToken                string `json:"TrustToken"`
		GuestSessionTimeout       string `json:"GuestSessionTimeout"`
		SessionTimeout            string `json:"SessionTimeout"`
		SessionTimeoutWarning     string `json:"SessionTimeoutWarning"`
		SessionTimeoutWLMethod    string `json:"SessionTimeoutWLMethod"`
		SessionTimeoutWLCommand   string `json:"SessionTimeoutWLCommand"`
		SingleSignOn              bool   `json:"SingleSignOn"`
		UserSpec                  bool   `json:"UserSpec"`
		SubUserSpec               bool   `json:"SubUserSpec"`
		UserSpecSource            bool   `json:"UserSpecSource"`
		ProtectedVirtualDirectory bool   `json:"ProtectedVirtualDirectory"`
		IntegratedDomainAuth      bool   `json:"IntegratedDomainAuth"`
		SiebEntSecToken           string `json:"SiebEntSecToken"`
		AuthenticationType        string `json:"AuthenticationType"`
		OAuthEndPoint             string `json:"OAuthEndPoint"`
		SessKeepAlive             string `json:"SessKeepAlive"`
	}

	// Defaults struct
	Defaults struct {
		AuthenticationProperties AuthenticationProperties `json:"AuthenticationProperties"`
		StatsPage                string                   `json:"StatsPage"`
		HTTPPort                 string                   `json:"HTTPPort"`
		HTTPSPort                string                   `json:"HTTPSPort"`
		EnableFQDN               bool                     `json:"EnableFQDN"`
		FQDN                     string                   `json:"FQDN"`
		DoCompression            string                   `json:"DoCompression"`
		Enabled                  bool                     `json:"Enabled"`
		SessionTracking          string                   `json:"SessionTracking"`
	}

	// LogProperties struct
	LogProperties struct {
		LogFILE           string `json:"LogFILE"`
		LogFileDirectory  string `json:"LogFileDirectory"`
		LogFileName       string `json:"LogFileName"`
		LogLevel          string `json:"LogLevel"`
		MaxLogFileSize    string `json:"MaxLogFileSize"`
		MaxLogBackupIndex string `json:"MaxLogBackupIndex"`
		Lang              string `json:"Lang"`
	}

	// ChannelDefault struct
	ChannelDefault struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// UI struct
	UI struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// EAI struct
	EAI struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// RESTInBound struct
	RESTInBound struct {
		AuthenticationProperties AuthenticationProperties `json:"AuthenticationProperties"`
		LogProperties            LogProperties            `json:"LogProperties"`
		MaxPoolSize              string                   `json:"MaxPoolSize"`
		MinPoolSize              string                   `json:"MinPoolSize"`
		Baseuri                  string                   `json:"Baseuri"`
	}

	// RESTOUTBOUND struct
	RESTOUTBOUND struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// SOAPOUTBOUND struct
	SOAPOUTBOUND struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// JBS struct
	JBS struct {
		LogProperties LogProperties `json:"LogProperties"`
		SessKeepAlive string        `json:"SessKeepAlive"`
	}

	// DAV struct
	DAV struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// Application struct
	Application struct {
		Name                     string                   `json:"Name"`
		Language                 string                   `json:"Language"`
		EnableExtServiceOnly     bool                     `json:"EnableExtServiceOnly"`
		UseAnonPool              bool                     `json:"UseAnonymousPool"`
		AnonUserPool             string                   `json:"AnonymousUserPool"`
		StartCommand             string                   `json:"StartCommand"`
		AuthenticationProperties AuthenticationProperties `json:"AuthenticationProperties"`
	}

	// SWSMConfigParams struct
	SWSMConfigParams struct {
		SWE            SWE            `json:"SWE"`
		ConnMgmt       ConnMgmt       `json:"ConnMgmt"`
		Defaults       Defaults       `json:"defaults"`
		ChannelDefault ChannelDefault `json:"CHANNELDEFAULT"`
		UI             UI             `json:"UI"`
		EAI            EAI            `json:"EAI"`
		RESTInBound    RESTInBound    `json:"RESTINBOUND"`
		RESTOUTBOUND   RESTOUTBOUND   `json:"RESTOUTBOUND"`
		SOAPOUTBOUND   SOAPOUTBOUND   `json:"SOAPOUTBOUND"`
		JBS            JBS            `json:"JBS"`
		DAV            DAV            `json:"DAV"`
		Applications   []Application  `json:"Applications"`
	}

	// SWSMProfile struct
	SWSMProfile struct {
		Profile          Profile          `json:"Profile"`
		SWSMConfigParams SWSMConfigParams `json:"ConfigParam"`
	}
)
SWSM 1.0 */

// SWSM 2.0
type (
	// SWE struct
	SWE struct {
		SWSMMode             string `json:"SWSMMode"`
		Language             string `json:"Language"`
		SessionMonitor       string `json:"SessionMonitor"`
		AllowStats           string `json:"AllowStats"`
		ClientRootDir        string `json:"ClientRootDir"`
		MaxQueryStringLength string `json:"MaxQueryStringLength"`
		SeedFile             string `json:"SeedFile"`
	}

	// ConnMgmt struct
	ConnMgmt struct {
		EnableVirtualHosts string `json:"EnableVirtualHosts"`
		VirtualHostsFile   string `json:"VirtualHostsFile"`
		CertFileName       string `json:"CertFileName"`
		CACertFileName     string `json:"CACertFileName"`
		KeyFileName        string `json:"KeyFileName"`
		KeyFilePassword    string `json:"KeyFilePassword"`
		PeerAuth           string `json:"PeerAuth"`
		PeerCertValidation string `json:"PeerCertValidation"`
	}

	// AuthenticationProperties struct
	AuthenticationProperties struct {
		EncryptedPassword         bool   `json:"EncryptedPassword"`
		EncryptSessionID          bool   `json:"EncryptSessionId"`
		AnonUserName              string `json:"AnonUserName"`
		AnonPassword              string `json:"AnonPassword"`
		TrustToken                string `json:"TrustToken"`
		GuestSessionTimeout       string `json:"GuestSessionTimeout"`
		SessionTimeout            string `json:"SessionTimeout"`
		SessionTimeoutWarning     string `json:"SessionTimeoutWarning"`
		SessionTimeoutWLMethod    string `json:"SessionTimeoutWLMethod"`
		SessionTimeoutWLCommand   string `json:"SessionTimeoutWLCommand"`
		SingleSignOn              bool   `json:"SingleSignOn"`
		UserSpec                  string `json:"UserSpec"`
		ProtectedVirtualDirectory string `json:"ProtectedVirtualDirectory"`
		SiebEntSecToken           string `json:"SiebEntSecToken"`
		AuthenticationType        string `json:"AuthenticationType"`
		OAuthEndPoint             string `json:"OAuthEndPoint"`
		SessKeepAlive             string `json:"SessKeepAlive"`
	}

	// Defaults struct
	Defaults struct {
		AuthenticationProperties AuthenticationProperties `json:"AuthenticationProperties"`
		StatsPage                string                   `json:"StatsPage"`
		HTTPPort                 string                   `json:"HTTPPort"`
		HTTPSPort                string                   `json:"HTTPSPort"`
		EnableFQDN               bool                     `json:"EnableFQDN"`
		FQDN                     string                   `json:"FQDN"`
		DoCompression            string                   `json:"DoCompression"`
		Enabled                  bool                     `json:"Enabled"`
		SessionTracking          string                   `json:"SessionTracking"`
	}

	// LogProperties struct
	LogProperties struct {
		LogFILE           string `json:"LogFILE"`
		LogFileDirectory  string `json:"LogFileDirectory"`
		LogFileName       string `json:"LogFileName"`
		LogLevel          string `json:"LogLevel"`
		MaxLogFileSize    string `json:"MaxLogFileSize"`
		MaxLogBackupIndex string `json:"MaxLogBackupIndex"`
		Lang              string `json:"Lang"`
	}

	// ChannelDefault struct
	ChannelDefault struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// UI struct
	UI struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// RESTResourceParam struct
	RESTResourceParam struct {
		Name         string `json:"Name"`
		Alias        string `json:"Alias"`
		DefaultValue string `json:"DefaultValue"`
	}

	// RESTResourceParamList struct
	RESTResourceParamList struct {
		OperationName string              `json:"OperationName"`
		ParamList     []RESTResourceParam `json:"ParamList"`
	}

	// RESTInBoundDefault struct
	RESTInBoundDefault struct {
		ObjMgr                   string                   `json:"ObjMgr"`
		MaxConnections           string                   `json:"MaxConnections"`
		MinConnections           string                   `json:"MinConnections"`
		Baseuri                  string                   `json:"Baseuri"`
		AuthenticationProperties AuthenticationProperties `json:"AuthenticationProperties"`
		LogProperties            LogProperties            `json:"LogProperties"`
		RESTResourceParamList    []RESTResourceParamList  `json:"RESTResourceParamList"`
	}

	// EAI struct
	EAI struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// RESTInBound struct
	RESTInBound struct {
		ObjMgr                   string                   `json:"ObjMgr"`
		Version                  string                   `json:"Version"`
		ResourceType             string                   `json:"ResourceType"`
		AuthenticationProperties AuthenticationProperties `json:"AuthenticationProperties"`
		LogProperties            LogProperties            `json:"LogProperties"`
		RESTResourceParamList    []RESTResourceParamList  `json:"RESTResourceParamList"`
	}

	// RESTOUTBOUND struct
	RESTOUTBOUND struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// SOAPOUTBOUND struct
	SOAPOUTBOUND struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// JBS struct
	JBS struct {
		LogProperties LogProperties `json:"LogProperties"`
		SessKeepAlive string        `json:"SessKeepAlive"`
	}

	// DAV struct
	DAV struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// Application struct
	Application struct {
		Name                     string                   `json:"Name"`
		Language                 string                   `json:"Language"`
		EnableExtServiceOnly     bool                     `json:"EnableExtServiceOnly"`
		UseAnonPool              bool                     `json:"UseAnonPool"`
		AnonUserPool             string                   `json:"AnonUserPool"`
		StartCommand             string                   `json:"StartCommand"`
		WebPublicRootDir         string                   `json:"WebPublicRootDir"`
		AuthenticationProperties AuthenticationProperties `json:"AuthenticationProperties"`
	}

	// SWSMConfigParams struct
	SWSMConfigParams struct {
		SWE                SWE                `json:"swe"`
		ConnMgmt           ConnMgmt           `json:"ConnMgmt"`
		Defaults           Defaults           `json:"defaults"`
		ChannelDefault     ChannelDefault     `json:"ChannelDefault"`
		UI                 UI                 `json:"UI"`
		RESTInBoundDefault RESTInBoundDefault `json:"RESTInBoundDefault"`
		EAI                EAI                `json:"EAI"`
		RESTInBound        []RESTInBound      `json:"RESTInBound"`
		RESTOUTBOUND       RESTOUTBOUND       `json:"RESTOutBound"`
		SOAPOUTBOUND       SOAPOUTBOUND       `json:"SOAPOutBound"`
		JBS                JBS                `json:"JBS"`
		DAV                DAV                `json:"DAV"`
		Applications       []Application      `json:"Applications"`
	}

	// SWSMProfile struct
	SWSMProfile struct {
		Profile          Profile          `json:"Profile"`
		SWSMConfigParams SWSMConfigParams `json:"ConfigParam"`
	}
)
