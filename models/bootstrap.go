package models

type (
	// BootstrapCG struct
	BootstrapCG struct {
		RegistryPort     string `json:"registryPort"`
		RegistryUsername string `json:"registryUserName"`
		RegistryPassword string `json:"registryPassword"`
		SecurityProfile  string `json:"SecurityProfile"`
		PrimaryLanguage  string `json:"PrimaryLanguage"`
	}

	// CGInfo struct
	CGInfo struct {
		CGHostURI string `json:"CGHostURI"`
	}
)
