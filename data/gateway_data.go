package data

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"../restapis"
)

// writeDataToFile -
func writeDataToFile(filename string, data interface{}) {
	// Encode data
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	if err != nil {
		log.Println("Encode:", err)
		return
	}

	// Write data
	err = ioutil.WriteFile(filename, buffer.Bytes(), os.ModePerm)
	if err != nil {
		log.Println("Write File:", err)
		// } else {
		// 	log.Println("Successfully write file:", filename)
	}
}

// SaveData -
func SaveData() {
	log.Printf("Saving Data...")

	// Create data directory
	os.Mkdir("GatewayData", os.ModePerm)

	writeDataToFile(filepath.Clean("GatewayData/CacheClientProfiles.data"), restapis.CacheClientProfiles)
	writeDataToFile(filepath.Clean("GatewayData/CacheServerProfiles.data"), restapis.CacheServerProfiles)
	writeDataToFile(filepath.Clean("GatewayData/CacheServerDeployments.data"), restapis.CacheServerDeployments)
	writeDataToFile(filepath.Clean("GatewayData/EnterpriseProfiles.data"), restapis.EnterpriseProfiles)
	writeDataToFile(filepath.Clean("GatewayData/EnterpriseDeployments.data"), restapis.EnterpriseDeployments)
	writeDataToFile(filepath.Clean("GatewayData/MigrationProfiles.data"), restapis.MigrationProfiles)
	writeDataToFile(filepath.Clean("GatewayData/MigrationDeployments.data"), restapis.MigrationDeployments)
	writeDataToFile(filepath.Clean("GatewayData/SecurityProfiles.data"), restapis.SecurityProfiles)
	writeDataToFile(filepath.Clean("GatewayData/ServerProfiles.data"), restapis.ServerProfiles)
	writeDataToFile(filepath.Clean("GatewayData/ServerDeployments.data"), restapis.ServerDeployments)
	writeDataToFile(filepath.Clean("GatewayData/SWSMProfiles.data"), restapis.SWSMProfiles)
	writeDataToFile(filepath.Clean("GatewayData/SWSMDeployments.data"), restapis.SWSMDeployments)

	log.Println("Done.")
}

// readDataFromFile -
func readDataFromFile(filename string, data interface{}) {
	// Read data
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Open File:", err)
		return
	}

	buffer := bytes.NewBuffer(bs)

	// Decode data
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(data)
	if err != nil {
		log.Println("Decode Error:", err)
		// } else {
		// 	log.Println("Successfully decode:", filename, len(restapis.CacheServerProfiles))
	}
}

// ReadData -
func ReadData() {
	log.Printf("Loading Data...")

	readDataFromFile(filepath.Clean("GatewayData/CacheClientProfiles.data"), &restapis.CacheClientProfiles)
	readDataFromFile(filepath.Clean("GatewayData/CacheServerProfiles.data"), &restapis.CacheServerProfiles)
	readDataFromFile(filepath.Clean("GatewayData/CacheServerDeployments.data"), &restapis.CacheServerDeployments)
	readDataFromFile(filepath.Clean("GatewayData/EnterpriseProfiles.data"), &restapis.EnterpriseProfiles)
	readDataFromFile(filepath.Clean("GatewayData/EnterpriseDeployments.data"), &restapis.EnterpriseDeployments)
	readDataFromFile(filepath.Clean("GatewayData/MigrationProfiles.data"), &restapis.MigrationProfiles)
	readDataFromFile(filepath.Clean("GatewayData/MigrationDeployments.data"), &restapis.MigrationDeployments)
	readDataFromFile(filepath.Clean("GatewayData/SecurityProfiles.data"), &restapis.SecurityProfiles)
	readDataFromFile(filepath.Clean("GatewayData/ServerProfiles.data"), &restapis.ServerProfiles)
	readDataFromFile(filepath.Clean("GatewayData/ServerDeployments.data"), &restapis.ServerDeployments)
	readDataFromFile(filepath.Clean("GatewayData/SWSMProfiles.data"), &restapis.SWSMProfiles)
	readDataFromFile(filepath.Clean("GatewayData/SWSMDeployments.data"), &restapis.SWSMDeployments)

	log.Println("Done.")
}
