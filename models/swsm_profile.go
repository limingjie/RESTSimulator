package models

type (
	// SWE struct
	SWE struct {
		Language             string `json:"Language"`
		SessionMonitor       bool   `json:"SessionMonitor"`
		AllowStats           bool   `json:"AllowStats"`
		MaxQueryStringLength int    `json:"MaxQueryStringLength"`
		SeedFile             string `json:"SeedFile"`
	}

	// ConnMgmt struct
	ConnMgmt struct {
		CertFileName       string `json:"CertFileName"`
		CACertFileName     string `json:"CACertFileName"`
		KeyFileName        string `json:"KeyFileName"`
		KeyFilePassword    string `json:"KeyFilePassword"`
		PeerAuth           bool   `json:"PeerAuth"`
		PeerCertValidation bool   `json:"PeerCertValidation"`
	}

	// AuthenticationProperties struct
	AuthenticationProperties struct {
		AnonUserName            string `json:"AnonUserName"`
		AnonPassword            string `json:"AnonPassword"`
		TrustToken              string `json:"TrustToken"`
		GuestSessionTimeout     int    `json:"GuestSessionTimeout"`
		SessionTimeout          int    `json:"SessionTimeout"`
		SessionTokenTimeout     int    `json:"SessionTokenTimeout"`
		SessionTokenMaxAge      int    `json:"SessionTokenMaxAge"`
		SessionTimeoutWLMethod  string `json:"SessionTimeoutWLMethod"`
		SessionTimeoutWLCommand string `json:"SessionTimeoutWLCommand"`
		SingleSignOn            bool   `json:"SingleSignOn"`
		UserSpec                string `json:"UserSpec"`
	}

	// Defaults struct
	Defaults struct {
		AuthenticationProperties AuthenticationProperties `json:"AuthenticationProperties"`
		StatsPage                string                   `json:"StatsPage"`
		EnableFQDN               bool                     `json:"EnableFQDN"`
		FQDN                     string                   `json:"FQDN"`
		DoCompression            bool                     `json:"DoCompression"`
	}

	// LogProperties struct
	LogProperties struct {
		LogLevel string `json:"LogLevel"`
	}

	// UI struct
	UI struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// RESTAuthenticationProperties struct
	RESTAuthenticationProperties struct {
		AnonUserName        string `json:"AnonUserName"`
		AnonPassword        string `json:"AnonPassword"`
		SessKeepAlive       int    `json:"SessKeepAlive"`
		AuthenticationType  string `json:"AuthenticationType"`
		TrustToken          string `json:"TrustToken"`
		OAuthEndPoint       string `json:"OAuthEndPoint"`
		UserSpec            string `json:"UserSpec"`
		ValidateCertificate bool   `json:"ValidateCertificate"`
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

	// RESTInBound struct
	RESTInBound struct {
		ObjectManager                string                       `json:"ObjectManager"`
		MaxConnections               int                          `json:"MaxConnections"`
		MinConnections               int                          `json:"MinConnections"`
		Baseuri                      string                       `json:"Baseuri"`
		RESTAuthenticationProperties RESTAuthenticationProperties `json:"RESTAuthenticationProperties"`
		LogProperties                LogProperties                `json:"LogProperties"`
		RESTResourceParamList        []RESTResourceParamList      `json:"RESTResourceParamList"`
	}

	// EAI struct
	EAI struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// RESTInBoundResource struct
	RESTInBoundResource struct {
		ResourceType          string                  `json:"ResourceType"`
		RESTResourceParamList []RESTResourceParamList `json:"RESTResourceParamList"`
	}

	// RESTOutBound struct
	RESTOutBound struct {
		LogProperties LogProperties `json:"LogProperties"`
	}

	// SOAPOutBound struct
	SOAPOutBound struct {
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
		ObjectManager            string                   `json:"ObjectManager"`
		AvailableInSiebelMobile  bool                     `json:"AvailableInSiebelMobile"`
		AppDisplayOrder          int                      `json:"AppDisplayOrder"`
		AppDisplayName           string                   `json:"AppDisplayName"`
		AppIcon                  string                   `json:"AppIcon"`
		EnableExtServiceOnly     bool                     `json:"EnableExtServiceOnly"`
		UseAnonPool              bool                     `json:"UseAnonPool"`
		AnonUserPool             int                      `json:"AnonUserPool"`
		StartCommand             string                   `json:"StartCommand"`
		EAISOAPNoSessInPref      bool                     `json:"EAISOAPNoSessInPref"`
		EAISOAPMaxRetry          int                      `json:"EAISOAPMaxRetry"`
		AuthenticationProperties AuthenticationProperties `json:"AuthenticationProperties"`
	}

	GatewayIdentity struct {
		GatewayHost string `json:"GatewayHost"`
		GatewayPort string `json:"GatewayPort"`
		AuthToken   string `json:"AuthToken"`
	}

	// SWSMConfigParams struct
	SWSMConfigParams struct {
		SWE                 SWE                   `json:"swe"`
		ConnMgmt            ConnMgmt              `json:"ConnMgmt"`
		Defaults            Defaults              `json:"defaults"`
		UI                  UI                    `json:"UI"`
		RESTInBound         RESTInBound           `json:"RESTInBound"`
		EAI                 EAI                   `json:"EAI"`
		RESTInBoundResource []RESTInBoundResource `json:"RESTInBoundResource"`
		RESTOutBound        RESTOutBound          `json:"RESTOutBound"`
		SOAPOutBound        SOAPOutBound          `json:"SOAPOutBound"`
		JBS                 JBS                   `json:"JBS"`
		DAV                 DAV                   `json:"DAV"`
		Applications        []Application         `json:"Applications"`
		GatewayIdentity     GatewayIdentity       `json:"GatewayIdentity"`
	}

	// SWSMProfile struct
	SWSMProfile struct {
		Profile          Profile          `json:"Profile"`
		SWSMConfigParams SWSMConfigParams `json:"ConfigParam"`
	}
)
