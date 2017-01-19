package models

type (
	// BootstrapCG struct
	BootstrapCG struct {
		RegistryPort     string `json:"registryPort"`
		RegistryUsername string `json:"registryUserName"`
		RegistryPassword string `json:"registryPassword"`
	}

	// CGInfo struct
	CGInfo struct {
		CGHostURI string `json:"CGHostURI"`
	}
)
