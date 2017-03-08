package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"../restapis"
)

// writeDataToFile -
func writeDataToFile(filename string, data interface{}) {
	// Encode data
	buffer, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Encode:", err)
		return
	}

	// Write data
	err = ioutil.WriteFile(filename, buffer, os.ModePerm)
	if err != nil {
		log.Println("Write File:", err)
		//} else {
		//	log.Println("Successfully write file:", filename)
	}
}

// SaveData -
func SaveData() {
	log.Println("Saving Data...")

	// Create data directory
	os.Mkdir("GatewayData", os.ModePerm)

	writeDataToFile(filepath.Clean("GatewayData/CacheClientProfiles.json"), restapis.CacheClientProfiles)
	writeDataToFile(filepath.Clean("GatewayData/CacheServerProfiles.json"), restapis.CacheServerProfiles)
	writeDataToFile(filepath.Clean("GatewayData/CacheServerDeployments.json"), restapis.CacheServerDeployments)
	writeDataToFile(filepath.Clean("GatewayData/ConstraintEngineProfiles.json"), restapis.ConstraintEngineProfiles)
	writeDataToFile(filepath.Clean("GatewayData/ConstraintEngineDeployments.json"), restapis.ConstraintEngineDeployments)
	writeDataToFile(filepath.Clean("GatewayData/EnterpriseProfiles.json"), restapis.EnterpriseProfiles)
	writeDataToFile(filepath.Clean("GatewayData/EnterpriseDeployments.json"), restapis.EnterpriseDeployments)
	writeDataToFile(filepath.Clean("GatewayData/MigrationProfiles.json"), restapis.MigrationProfiles)
	writeDataToFile(filepath.Clean("GatewayData/MigrationDeployments.json"), restapis.MigrationDeployments)
	writeDataToFile(filepath.Clean("GatewayData/SecurityProfiles.json"), restapis.SecurityProfiles)
	writeDataToFile(filepath.Clean("GatewayData/ServerProfiles.json"), restapis.ServerProfiles)
	writeDataToFile(filepath.Clean("GatewayData/ServerDeployments.json"), restapis.ServerDeployments)
	writeDataToFile(filepath.Clean("GatewayData/SWSMProfiles.json"), restapis.SWSMProfiles)
	writeDataToFile(filepath.Clean("GatewayData/SWSMDeployments.json"), restapis.SWSMDeployments)

	log.Println("Done.")
}

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

// LoadData -
func LoadData() {
	log.Println("Loading Data...")

	loadDataFromFile(filepath.Clean("GatewayData/CacheClientProfiles.json"), &restapis.CacheClientProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/CacheServerProfiles.json"), &restapis.CacheServerProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/CacheServerDeployments.json"), &restapis.CacheServerDeployments)
	loadDataFromFile(filepath.Clean("GatewayData/ConstraintEngineProfiles.json"), &restapis.ConstraintEngineProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/ConstraintEngineDeployments.json"), &restapis.ConstraintEngineDeployments)
	loadDataFromFile(filepath.Clean("GatewayData/EnterpriseProfiles.json"), &restapis.EnterpriseProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/EnterpriseDeployments.json"), &restapis.EnterpriseDeployments)
	loadDataFromFile(filepath.Clean("GatewayData/MigrationProfiles.json"), &restapis.MigrationProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/MigrationDeployments.json"), &restapis.MigrationDeployments)
	loadDataFromFile(filepath.Clean("GatewayData/SecurityProfiles.json"), &restapis.SecurityProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/ServerProfiles.json"), &restapis.ServerProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/ServerDeployments.json"), &restapis.ServerDeployments)
	loadDataFromFile(filepath.Clean("GatewayData/SWSMProfiles.json"), &restapis.SWSMProfiles)
	loadDataFromFile(filepath.Clean("GatewayData/SWSMDeployments.json"), &restapis.SWSMDeployments)

	log.Println("Done.")
}
