package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"../models"
	"github.com/julienschmidt/httprouter"
)

// EnterpriseDeployments map
var EnterpriseDeployments = make(map[string]models.EnterpriseDeployment)

func saveEnterpriseDeployments() {
	WriteFile(filepath.Clean("GatewayData/EnterpriseDeployments.json"), EnterpriseDeployments)
}

func deployEnterpriseProfile(profileName string) {
	profile, ok := EnterpriseProfiles[profileName]
	if ok {
		profile.Profile.Deploy()
		EnterpriseProfiles[profileName] = profile
	}
}

func undeployEnterpriseProfile(profileName string) {
	profile, ok := EnterpriseProfiles[profileName]
	if ok {
		profile.Profile.Undeploy()
		EnterpriseProfiles[profileName] = profile
	}
}

// PostEnterpriseDeployment - POST /deployments/enterprises
func PostEnterpriseDeployment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deployment := models.EnterpriseDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	log.Println(r.RemoteAddr, "PostEnterpriseDeployment", string(msg))

	_, ok := EnterpriseProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Inexist Deployment Profile.")
		return
	}

	deploymentName := deployment.EnterpriseDeployParams.SiebelEnterprise
	if len(deploymentName) > 0 {
		_, ok := EnterpriseDeployments[deploymentName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Enterprise deployment with same name already exists.")
		} else {
			deployment.Deployment.Check()
			EnterpriseDeployments[deploymentName] = deployment

			deployEnterpriseProfile(deployment.Deployment.ProfileName)

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")

			saveEnterpriseDeployments()
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Name.")
	}
}

// GetEnterpriseDeployments - GET /deployments/enterprises
func GetEnterpriseDeployments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var deploymentsJSON bytes.Buffer
	deploymentsJSON.WriteString("{\"EnterpriseDeployment\":[")

	log.Println(r.RemoteAddr, "GetEnterpriseDeployments", strconv.Itoa(len(EnterpriseDeployments)))

	ok := false
	for _, deployment := range EnterpriseDeployments {
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

// GetEnterpriseDeployment - GET /deployments/enterprises/:name
func GetEnterpriseDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment, ok := EnterpriseDeployments[ps.ByName("name")]

	log.Println(r.RemoteAddr, "GetEnterpriseDeployment", ps.ByName("name"))

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		deploymentJSON, _ := json.Marshal(deployment)
		fmt.Fprintf(w, "%s", deploymentJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise deployment does not exist.")
	}
}

// PutEnterpriseDeployment - PUT /deployments/enterprises/:name
func PutEnterpriseDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment := models.EnterpriseDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	log.Println(r.RemoteAddr, "PutEnterpriseDeployment", string(msg))

	_, ok := EnterpriseProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Profile.")
		return
	}

	deploymentName := ps.ByName("name")
	_, ok = EnterpriseDeployments[deploymentName]
	if ok {
		if deployment.EnterpriseDeployParams.SiebelEnterprise == deploymentName {
			oldProfileName := EnterpriseDeployments[deploymentName].Deployment.ProfileName
			if deployment.Deployment.ProfileName != oldProfileName {
				deployEnterpriseProfile(deployment.Deployment.ProfileName)
				undeployEnterpriseProfile(oldProfileName)
			}

			deployment.Deployment.Check()
			EnterpriseDeployments[deploymentName] = deployment

			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")

			saveEnterpriseDeployments()
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Enterprise deployment name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise deployment does not exist.")
	}
}

// DeleteEnterpriseDeployment - DELETE /deployments/enterprises/:name
func DeleteEnterpriseDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deploymentName := ps.ByName("name")

	log.Println(r.RemoteAddr, "DeleteEnterpriseDeployment", deploymentName)

	deployment, ok := EnterpriseDeployments[deploymentName]
	if ok {
		_, ok := EnterpriseProfiles[deployment.Deployment.ProfileName]
		if ok {
			undeployEnterpriseProfile(deployment.Deployment.ProfileName)
		}

		delete(EnterpriseDeployments, deploymentName)

		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")

		saveEnterpriseDeployments()
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise deployment does not exist.")
	}
}
