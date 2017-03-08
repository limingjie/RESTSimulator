package models

type (
	// TangoCoherenceOverride struct
	TangoCoherenceOverride struct {
		TangoConfig string `json:"tangoConfig"`
		ClusterName string `json:"clusterName"`
	}

	// CachingSchemeMapping struct
	CachingSchemeMapping struct {
		CacheName  string `json:"cacheName"`
		SchemeName string `json:"schemeName"`
	}

	// ProxyScheme struct
	ProxyScheme struct {
		ServiceName    string `json:"serviceName"`
		ThreadCountMax int    `json:"threadCountMax"`
		ThreadCountMin int    `json:"threadCountMin"`
		Address        string `json:"address"`
		Port           int    `json:"port"`
		Enabled        bool   `json:"enabled"`
		Autostart      bool   `json:"autostart"`
	}

	// DistributedScheme struct
	DistributedScheme struct {
		SchemeName       string `json:"schemeName"`
		SchemeRef        string `json:"schemeRef"`
		ServiceName      string `json:"serviceName"`
		Autostart        bool   `json:"autostart"`
		BackingMapScheme string `json:"backingMapScheme"`
	}

	// LocalScheme struct
	LocalScheme struct {
		SchemeName     string `json:"schemeName"`
		EvictionPolicy string `json:"evictionPolicy"`
		HighUnits      int    `json:"highUnits"`
		UnitCalculator string `json:"unitCalculator"`
		UnitFactor     int    `json:"unitFactor"`
		ExpiryDelay    string `json:"expiryDelay"`
	}

	// CachingSchemes struct
	CachingSchemes struct {
		ProxyScheme       ProxyScheme         `json:"proxyScheme"`
		DistributedScheme []DistributedScheme `json:"distributedScheme"`
		LocalScheme       []LocalScheme       `json:"localScheme"`
	}

	// CacheConfigObj struct
	CacheConfigObj struct {
		CachingSchemeMapping []CachingSchemeMapping `json:"cachingSchemeMapping"`
		CachingSchemes       CachingSchemes         `json:"cachingSchemes"`
	}

	// CacheConfig struct
	CacheConfig struct {
		CacheConfigXML string         `json:"cacheConfigXml"`
		CacheConfigObj CacheConfigObj `json:"cacheConfigObj"`
	}

	// ServerConfigParam struct
	ServerConfigParam struct {
		TangoCoherenceOverride TangoCoherenceOverride `json:"tangoCoherenceOverride"`
		CacheConfig            CacheConfig            `json:"cacheConfig"`
	}

	// CacheConfigParams struct
	CacheConfigParams struct {
		LogProperties     LogProperties     `json:"LogProperties"`
		SecurityProfile   string            `json:"SecurityProfile"`
		ServerConfigParam ServerConfigParam `json:"ServerConfigParam"`
		GatewayIdentity   GatewayIdentity   `json:"GatewayIdentity"`
	}

	// CacheServerProfile struct
	CacheServerProfile struct {
		Profile           Profile           `json:"Profile"`
		CacheConfigParams CacheConfigParams `json:"CacheConfigParams"`
	}
)
