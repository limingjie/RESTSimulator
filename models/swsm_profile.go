package models

type (
	// SWE struct
	SWE struct {
		Language             string `json:"language"`
		SessionMonitor       string `json:"sessionMonitor"`
		AllowStats           string `json:"allowStats"`
		ClientRootDir        string `json:"clientRootDir"`
		MaxQueryStringLength string `json:"maxQueryStringLength"`
	}

	// ConnMgmt struct
	ConnMgmt struct {
		EnableVirtualHosts      string `json:"enableVirtualHosts"`
		VirtualHostsFileContent string `json:"virtualHostsFileContent"`
		CertFileName            string `json:"certFileName"`
		CACertFileName          string `json:"cACertFileName"`
		KeyFileName             string `json:"keyFileName"`
		KeyFilePassword         string `json:"keyFilePassword"`
		PeerAuth                string `json:"peerAuth"`
		PeerCertValidation      string `json:"peerCertValidation"`
	}

	// AuthenticationProperties struct
	AuthenticationProperties struct {
		EncryptedPassword         string `json:"encryptedPassword"`
		EncryptSessionID          string `json:"encryptSessionId"`
		AnonUserName              string `json:"anonUserName"`
		AnonPassword              string `json:"anonPassword"`
		ClientCertificate         string `json:"clientCertificate"`
		TrustToken                string `json:"trustToken"`
		GuestSessionTimeout       string `json:"guestSessionTimeout"`
		SessionTimeout            string `json:"sessionTimeout"`
		SessionTimeoutWarning     string `json:"sessionTimeoutWarning"`
		SessionTimeoutWLMethod    string `json:"sessionTimeoutWLMethod"`
		SessionTimeoutWLCommand   string `json:"sessionTimeoutWLCommand"`
		SingleSignOn              string `json:"singleSignOn"`
		UserSpec                  string `json:"userSpec"`
		SubUserSpec               string `json:"subUserSpec"`
		UserSpecSource            string `json:"userSpecSource"`
		ProtectedVirtualDirectory string `json:"protectedVirtualDirectory"`
		IntegratedDomainAuth      string `json:"integratedDomainAuth"`
		SiebEntSecToken           string `json:"siebEntSecToken"`
		AuthenticationType        string `json:"authenticationType"`
		OAuthEndPoint             string `json:"oAuthEndPoint"`
		SessKeepAlive             string `json:"sessKeepAlive"`
	}

	// Defaults struct
	Defaults struct {
		AuthenticationProperties AuthenticationProperties `json:"authenticationProperties"`
		StatsPage                string                   `json:"statsPage"`
		HTTPPort                 string                   `json:"httpPort"`
		HTTPSPort                string                   `json:"httpSPort"`
		EnableFQDN               string                   `json:"enableFQDN"`
		FQDN                     string                   `json:"fqdn"`
		DoCompression            string                   `json:"doCompression"`
		Enabled                  string                   `json:"enabled"`
		SessionTracking          string                   `json:"sessionTracking"`
	}

	// LogProperties struct
	LogProperties struct {
		LogFILE           string `json:"logFILE"`
		LogFileDirectory  string `json:"logFileDirectory"`
		LogFileName       string `json:"logFileName"`
		LogLevel          string `json:"logLevel"`
		MaxLogFileSize    string `json:"maxLogFileSize"`
		MaxLogBackupIndex string `json:"maxLogBackupIndex"`
		Lang              string `json:"lang"`
	}

	// CHANNELDEFAULT struct
	CHANNELDEFAULT struct {
		LogProperties LogProperties `json:"logProperties"`
	}

	// UI struct
	UI struct {
		LogProperties LogProperties `json:"logProperties"`
	}

	// EAI struct
	EAI struct {
		LogProperties LogProperties `json:"logProperties"`
	}

	// RESTINBOUND struct
	RESTINBOUND struct {
		AuthenticationProperties AuthenticationProperties `json:"authenticationProperties"`
		LogProperties            LogProperties            `json:"logProperties"`
		MaxPoolSize              string                   `json:"maxPoolSize"`
		MinPoolSize              string                   `json:"minPoolSize"`
		Baseuri                  string                   `json:"baseuri"`
	}

	// RESTOUTBOUND struct
	RESTOUTBOUND struct {
		LogProperties LogProperties `json:"logProperties"`
	}

	// SOAPOUTBOUND struct
	SOAPOUTBOUND struct {
		LogProperties LogProperties `json:"logProperties"`
	}

	// JBS struct
	JBS struct {
		LogProperties LogProperties `json:"logProperties"`
		SessKeepAlive string        `json:"sessKeepAlive"`
	}

	// DAV struct
	DAV struct {
		LogProperties LogProperties `json:"logProperties"`
	}

	// Application struct
	Application struct {
		Name                     string                   `json:"name"`
		Language                 string                   `json:"language"`
		EnableExtServiceOnly     string                   `json:"enableExtServiceOnly"`
		UseAnonPool              string                   `json:"useAnonPool"`
		AnonUserPool             string                   `json:"anonUserPool"`
		StartCommand             string                   `json:"startCommand"`
		AuthenticationProperties AuthenticationProperties `json:"authenticationProperties"`
	}

	// SWSMConfigParams struct
	SWSMConfigParams struct {
		SWE                      SWE                      `json:"swe"`
		ConnMgmt                 ConnMgmt                 `json:"connMgmt"`
		AuthenticationProperties AuthenticationProperties `json:"authenticationProperties"`
		Defaults                 Defaults                 `json:"defaults"`
		LogProperties            LogProperties            `json:"logProperties"`
		CHANNELDEFAULT           CHANNELDEFAULT           `json:"channelDefault"`
		UI                       UI                       `json:"ui"`
		EAI                      EAI                      `json:"eai"`
		RESTINBOUND              RESTINBOUND              `json:"restInbound"`
		RESTOUTBOUND             RESTOUTBOUND             `json:"restOutbound"`
		SOAPOUTBOUND             SOAPOUTBOUND             `json:"soapOutbound"`
		JBS                      JBS                      `json:"jbs"`
		DAV                      DAV                      `json:"dav"`
		Applications             []Application            `json:"applications"`
	}

	// SWSMProfile struct
	SWSMProfile struct {
		Profile          Profile          `json:"profile"`
		SWSMConfigParams SWSMConfigParams `json:"swsmConfigParams"`
	}
)
