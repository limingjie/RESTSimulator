package models

type (
	// Use CachingSchemeMapping of CacheServer Profile

	// SchemeParam struct
	SchemeParam struct {
		SchemeName  string `json:"schemeName"`
		ServiceName string `json:"serviceName"`
	}

	// TCPInitiator struct
	TCPInitiator struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	}

	// OutgoingMessageHandler struct
	OutgoingMessageHandler struct {
		HeartbeatInterval string `json:"heartbeatInterval"`
		HeartbeatTimeout  string `json:"heartbeatTimeout"`
		RequestTimeout    string `json:"requestTimeout"`
	}

	// InitiatorConfig struct
	InitiatorConfig struct {
		TCPInitiator           TCPInitiator           `json:"tcpInitiator"`
		OutgoingMessageHandler OutgoingMessageHandler `json:"outgoingMessageHandler"`
	}

	// CacheSchemes struct
	CacheSchemes struct {
		SchemeParam     SchemeParam     `json:"schemeParam"`
		InitiatorConfig InitiatorConfig `json:"initiatorConfig"`
	}

	// CacheClientConfig struct
	CacheClientConfig struct {
		CacheConfigXML       string               `json:"cacheConfigXml"`
		CachingSchemeMapping CachingSchemeMapping `json:"cachingSchemeMapping"`
		CacheSchemes         CacheSchemes         `json:"cacheSchemes"`
	}

	// ClientConfigParam struct
	ClientConfigParam struct {
		CacheClientConfig CacheClientConfig `json:"cacheConfig"`
	}

	// CacheClientConfigParams struct
	CacheClientConfigParams struct {
		LogProperties     LogProperties     `json:"LogProperties"`
		ClientConfigParam ClientConfigParam `json:"ClientConfigParam"`
	}

	// CacheClientProfile struct
	CacheClientProfile struct {
		Profile                 Profile                 `json:"Profile"`
		CacheClientConfigParams CacheClientConfigParams `json:"CacheConfigParams"`
	}
)
