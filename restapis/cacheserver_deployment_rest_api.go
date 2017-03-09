package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../models"
	"github.com/julienschmidt/httprouter"
)

// CacheServerDeployments map
var CacheServerDeployments = make(map[string]models.CacheServerDeployment)

func deployCacheServerProfile(profileName string) {
	profile, ok := CacheServerProfiles[profileName]
	if ok {
		profile.Profile.Deploy()
		CacheServerProfiles[profileName] = profile
	}
}

func undeployCacheServerProfile(profileName string) {
	profile, ok := CacheServerProfiles[profileName]
	if ok {
		profile.Profile.Undeploy()
		CacheServerProfiles[profileName] = profile
	}
}

// PostCacheServerDeployment - POST /deployments/cacheserver
func PostCacheServerDeployment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deployment := models.CacheServerDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	log.Println("PostCacheServerDeployment", string(msg))

	_, ok := CacheServerProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Inexist Deployment Profile.")
		return
	}

	deploymentName := deployment.CacheServerDeployParams.CacheServerAgentNode
	if len(deploymentName) > 0 {
		_, ok := CacheServerDeployments[deploymentName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: CacheServer deployment with same name already exists.")
		} else {
			deployment.Deployment.Check()
			CacheServerDeployments[deploymentName] = deployment

			deployCacheServerProfile(deployment.Deployment.ProfileName)

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Name.")
	}
}

// GetCacheServerDeployments - GET /deployments/cacheserver
func GetCacheServerDeployments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var deploymentsJSON bytes.Buffer
	deploymentsJSON.WriteString("{\"CacheServerDeployment\":[")

	log.Println("GetCacheServerDeployments", strconv.Itoa(len(CacheServerDeployments)))

	ok := false
	for _, deployment := range CacheServerDeployments {
		if ok {
			deploymentsJSON.WriteString(",")
		} else {
			ok = true
		}
		deploymentJSON, _ := json.Marshal(deployment)
		deploymentsJSON.Write(deploymentJSON)
	}

	deploymentsJSON.WriteString("]}")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", deploymentsJSON.String())
}

// GetCacheServerDeployment - GET /deployments/cacheserver/:name
func GetCacheServerDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment, ok := CacheServerDeployments[ps.ByName("name")]

	log.Println("GetCacheServerDeployment", ps.ByName("name"))

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		deploymentJSON, _ := json.Marshal(deployment)
		fmt.Fprintf(w, "%s", deploymentJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: CacheServer deployment does not exist.")
	}
}

// PutCacheServerDeployment - PUT /deployments/cacheserver/:name
func PutCacheServerDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment := models.CacheServerDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	log.Println("PutCacheServerDeployment", string(msg))

	_, ok := CacheServerProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Profile.")
		return
	}

	deploymentName := ps.ByName("name")
	_, ok = CacheServerDeployments[deploymentName]
	if ok {
		if deployment.CacheServerDeployParams.CacheServerAgentNode == deploymentName {
			oldProfileName := CacheServerDeployments[deploymentName].Deployment.ProfileName
			if deployment.Deployment.ProfileName != oldProfileName {
				deployCacheServerProfile(deployment.Deployment.ProfileName)
				undeployCacheServerProfile(oldProfileName)
			}

			CacheServerDeployments[deploymentName] = deployment

			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: CacheServer deployment name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: CacheServer deployment does not exist.")
	}
}

// DeleteCacheServerDeployment - DELETE /deployments/cacheserver/:name
func DeleteCacheServerDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deploymentName := ps.ByName("name")

	log.Println("DeleteCacheServerDeployment", deploymentName)

	deployment, ok := CacheServerDeployments[deploymentName]
	if ok {
		_, ok := CacheServerProfiles[deployment.Deployment.ProfileName]
		if ok {
			undeployCacheServerProfile(deployment.Deployment.ProfileName)
		}

		delete(CacheServerDeployments, deploymentName)

		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: CacheServer deployment does not exist.")
	}
}
