package restapis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../logger"
	"../models"
	"github.com/julienschmidt/httprouter"
)

// EnterpriseDeployments map
var EnterpriseDeployments = make(map[string]models.EnterpriseDeployment)

func deployEntpriseProfile(profileName string) {
	profile, ok := EnterpriseProfiles[profileName]
	if ok {
		profile.Profile.Deploy()
		EnterpriseProfiles[profileName] = profile
	}
}

func undeployEntpriseProfile(profileName string) {
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
	logger.Logger("PostEnterpriseDeployment", string(msg))

	_, ok := EnterpriseProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Inexist Deployment Profile.")
		return
	}

	deploymentName := deployment.EnterpriseDeployParams.EnterpriseServer
	if len(deploymentName) > 0 {
		_, ok := EnterpriseDeployments[deploymentName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Enterprise deployment with same name already exists.")
		} else {
			deployment.Deployment.Check()
			EnterpriseDeployments[deploymentName] = deployment

			deployEntpriseProfile(deployment.Deployment.ProfileName)

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Name.")
	}
}

// GetEnterpriseDeployments - GET /deployments/enterprises
func GetEnterpriseDeployments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var deploymentsJSON bytes.Buffer
	deploymentsJSON.WriteString("{\"enterpriseDeployment\":[")

	logger.Logger("GetEnterpriseDeployments", strconv.Itoa(len(EnterpriseDeployments)))

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

// GetEnterpriseDeployment - GET /deployments/enterprises/:deploymentname
func GetEnterpriseDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment, ok := EnterpriseDeployments[ps.ByName("deploymentname")]

	logger.Logger("GetEnterpriseDeployment", ps.ByName("deploymentname"))

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

// PutEnterpriseDeployment - PUT /deployments/enterprises/:deploymentname
func PutEnterpriseDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment := models.EnterpriseDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	logger.Logger("PutEnterpriseDeployment", string(msg))

	_, ok := EnterpriseProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Profile.")
		return
	}

	deploymentName := ps.ByName("deploymentname")
	_, ok = EnterpriseDeployments[deploymentName]
	if ok {
		if deployment.EnterpriseDeployParams.EnterpriseServer == deploymentName {
			oldProfileName := EnterpriseDeployments[deploymentName].Deployment.ProfileName
			if deployment.Deployment.ProfileName != oldProfileName {
				deployEntpriseProfile(deployment.Deployment.ProfileName)
				undeployEntpriseProfile(oldProfileName)
			}

			EnterpriseDeployments[deploymentName] = deployment

			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Enterprise deployment name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise deployment does not exist.")
	}
}

// DeleteEnterpriseDeployment - DELETE /deployments/enterprises/:deploymentname
func DeleteEnterpriseDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deploymentName := ps.ByName("deploymentname")

	logger.Logger("DeleteEnterpriseDeployment", deploymentName)

	deployment, ok := EnterpriseDeployments[deploymentName]
	if ok {
		_, ok := EnterpriseProfiles[deployment.Deployment.ProfileName]
		if ok {
			undeployEntpriseProfile(deployment.Deployment.ProfileName)
		}

		delete(EnterpriseDeployments, deploymentName)

		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise deployment does not exist.")
	}
}
