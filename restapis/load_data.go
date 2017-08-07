package restapis

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

// loadDataFromFile -
func loadDataFromFile(filename string, data interface{}) {
	// Read data
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		// log.Println("Open File:", err)
		return
	}

	// Decode data
	err = json.Unmarshal(buffer, data)
	if err != nil {
		log.Println("Decode Error:", err)
		//} else {
		//	log.Println("Successfully decode:", filename, len(restapis.CacheServerProfiles))
	}
}

// LoadAll -
func LoadAll() {
	log.Println("Loading Data...")

	loadDataFromFile(filepath.Clean("GatewayData/CacheClientProfiles.json"), &CacheClientProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/CacheServerProfiles.json"), &CacheServerProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/CacheServerDeployments.json"), &CacheServerDeployments)
	loadDataFromFile(filepath.Clean("GatewayData/ConstraintEngineProfiles.json"), &ConstraintEngineProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/ConstraintEngineDeployments.json"), &ConstraintEngineDeployments)
	loadDataFromFile(filepath.Clean("GatewayData/GatewayClusterProfiles.json"), &GatewayClusterProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/GatewayClusterDeployments.json"), &GatewayClusterDevelopments)
	loadDataFromFile(filepath.Clean("GatewayData/EnterpriseProfiles.json"), &EnterpriseProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/EnterpriseDeployments.json"), &EnterpriseDeployments)
	loadDataFromFile(filepath.Clean("GatewayData/MigrationProfiles.json"), &MigrationProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/MigrationDeployments.json"), &MigrationDeployments)
	loadDataFromFile(filepath.Clean("GatewayData/SecurityProfiles.json"), &SecurityProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/ServerProfiles.json"), &ServerProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/ServerDeployments.json"), &ServerDeployments)
	loadDataFromFile(filepath.Clean("GatewayData/SWSMProfiles.json"), &SWSMProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/SWSMDeployments.json"), &SWSMDeployments)

	log.Println("Done.")
}
