package models

type (
	// CacheConfigParams struct
	CacheConfigParams struct {
		TangoConfig string `json:"tango-config"`
		CacheConfig string `json:"cache-config"`
		PofConfig   string `json:"pof-config"`
	}

	// CacheProfile struct
	CacheProfile struct {
		Profile           Profile           `json:"Profile"`
		CacheConfigParams CacheConfigParams `json:"CacheConfigParams"`
	}
)
