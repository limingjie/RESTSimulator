package restapis

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// WriteFile -
func WriteFile(filename string, data interface{}) {
	// Create directories
	dir := filepath.Dir(filename)
	errMkdir := os.MkdirAll(dir, os.ModePerm)
	if errMkdir != nil {
		log.Println("Failed to create path", dir, errMkdir)
	}

	// Encode data
	buffer, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Failed to encode:", err)
		return
	}

	// Write data
	err = ioutil.WriteFile(filename, buffer, os.ModePerm)
	if err != nil {
		log.Println("Failed to write File:", err)
		//} else {
		//	log.Println("Successfully write file:", filename)
	}
}
