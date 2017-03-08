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

// SWSMDeployments map
var SWSMDeployments = make(map[string]models.SWSMDeployment)

func deploySWSMProfile(profileName string) {
	profile, ok := SWSMProfiles[profileName]
	if ok {
		profile.Profile.Deploy()
		SWSMProfiles[profileName] = profile
	}
}

func undeploySWSMProfile(profileName string) {
	profile, ok := SWSMProfiles[profileName]
	if ok {
		profile.Profile.Undeploy()
		SWSMProfiles[profileName] = profile
	}
}

// PostSWSMDeployment - POST /deployments/swsm
func PostSWSMDeployment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deployment := models.SWSMDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	log.Println("PostSWSMDeployment", string(msg))

	_, ok := SWSMProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Inexist Deployment Profile.")
		return
	}

	deploymentName := deployment.SWSMDeployParams.SWSMNode
	if len(deploymentName) > 0 {
		_, ok := SWSMDeployments[deploymentName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: SWSM deployment with same name already exists.")
		} else {
			deployment.Deployment.Check()
			SWSMDeployments[deploymentName] = deployment

			deploySWSMProfile(deployment.Deployment.ProfileName)

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Name.")
	}
}

// GetSWSMDeployments - GET /deployments/swsm
func GetSWSMDeployments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var deploymentsJSON bytes.Buffer
	deploymentsJSON.WriteString("{\"SWSMDeployment\":[")

	log.Println("GetSWSMDeployments", strconv.Itoa(len(SWSMDeployments)))

	ok := false
	for _, deployment := range SWSMDeployments {
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

// GetSWSMDeployment - GET /deployments/swsm/:deploymentname
func GetSWSMDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment, ok := SWSMDeployments[ps.ByName("deploymentname")]

	log.Println("GetSWSMDeployment", ps.ByName("deploymentname"))

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		deploymentJSON, _ := json.Marshal(deployment)
		fmt.Fprintf(w, "%s", deploymentJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: SWSM deployment does not exist.")
	}
}

// PutSWSMDeployment - PUT /deployments/swsm/:deploymentname
func PutSWSMDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment := models.SWSMDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	log.Println("PutSWSMDeployment", string(msg))

	_, ok := SWSMProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Profile.")
		return
	}

	deploymentName := ps.ByName("deploymentname")
	_, ok = SWSMDeployments[deploymentName]
	if ok {
		if deployment.SWSMDeployParams.SWSMNode == deploymentName {
			oldProfileName := SWSMDeployments[deploymentName].Deployment.ProfileName
			if deployment.Deployment.ProfileName != oldProfileName {
				deploySWSMProfile(deployment.Deployment.ProfileName)
				undeploySWSMProfile(oldProfileName)
			}

			SWSMDeployments[deploymentName] = deployment

			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: SWSM deployment name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: SWSM deployment does not exist.")
	}
}

// DeleteSWSMDeployment - DELETE /deployments/swsm/:deploymentname
func DeleteSWSMDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deploymentName := ps.ByName("deploymentname")

	log.Println("DeleteSWSMDeployment", deploymentName)

	deployment, ok := SWSMDeployments[deploymentName]
	if ok {
		_, ok := SWSMProfiles[deployment.Deployment.ProfileName]
		if ok {
			undeploySWSMProfile(deployment.Deployment.ProfileName)
		}

		delete(SWSMDeployments, deploymentName)

		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: SWSM deployment does not exist.")
	}
}
