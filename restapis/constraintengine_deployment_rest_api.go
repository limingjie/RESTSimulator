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

// ConstraintEngineDeployments map
var ConstraintEngineDeployments = make(map[string]models.ConstraintEngineDeployment)

func deployConstraintEngineProfile(profileName string) {
	profile, ok := ConstraintEngineProfiles[profileName]
	if ok {
		profile.Profile.Deploy()
		ConstraintEngineProfiles[profileName] = profile
	}
}

func undeployConstraintEngineProfile(profileName string) {
	profile, ok := ConstraintEngineProfiles[profileName]
	if ok {
		profile.Profile.Undeploy()
		ConstraintEngineProfiles[profileName] = profile
	}
}

// PostConstraintEngineDeployment - POST /deployments/constraintengines
func PostConstraintEngineDeployment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deployment := models.ConstraintEngineDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	log.Println("PostConstraintEngineDeployment", string(msg))

	_, ok := ConstraintEngineProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Inexist Deployment Profile.")
		return
	}

	deploymentName := deployment.ConstraintEngineDeployParams.DeploymentName
	if len(deploymentName) > 0 {
		_, ok := ConstraintEngineDeployments[deploymentName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: ConstraintEngine deployment with same name already exists.")
		} else {
			deployment.Deployment.Check()
			ConstraintEngineDeployments[deploymentName] = deployment

			deployConstraintEngineProfile(deployment.Deployment.ProfileName)

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Name.")
	}
}

// GetConstraintEngineDeployments - GET /deployments/constraintengines
func GetConstraintEngineDeployments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var deploymentsJSON bytes.Buffer
	deploymentsJSON.WriteString("{\"ConstraintEngineDeployment\":[")

	log.Println("GetConstraintEngineDeployments", strconv.Itoa(len(ConstraintEngineDeployments)))

	ok := false
	for _, deployment := range ConstraintEngineDeployments {
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

// GetConstraintEngineDeployment - GET /deployments/constraintengines/:deploymentname
func GetConstraintEngineDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment, ok := ConstraintEngineDeployments[ps.ByName("deploymentname")]

	log.Println("GetConstraintEngineDeployment", ps.ByName("deploymentname"))

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		deploymentJSON, _ := json.Marshal(deployment)
		fmt.Fprintf(w, "%s", deploymentJSON)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: ConstraintEngine deployment does not exist.")
	}
}

// PutConstraintEngineDeployment - PUT /deployments/constraintengines/:deploymentname
func PutConstraintEngineDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment := models.ConstraintEngineDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	log.Println("PutConstraintEngineDeployment", string(msg))

	_, ok := ConstraintEngineProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Profile.")
		return
	}

	deploymentName := ps.ByName("deploymentname")
	_, ok = ConstraintEngineDeployments[deploymentName]
	if ok {
		if deployment.ConstraintEngineDeployParams.DeploymentName == deploymentName {
			oldProfileName := ConstraintEngineDeployments[deploymentName].Deployment.ProfileName
			if deployment.Deployment.ProfileName != oldProfileName {
				deployConstraintEngineProfile(deployment.Deployment.ProfileName)
				undeployConstraintEngineProfile(oldProfileName)
			}

			ConstraintEngineDeployments[deploymentName] = deployment

			w.WriteHeader(200)
			fmt.Fprintf(w, "Succeed.")
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: ConstraintEngine deployment name does not match.")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: ConstraintEngine deployment does not exist.")
	}
}

// DeleteConstraintEngineDeployment - DELETE /deployments/constraintengines/:deploymentname
func DeleteConstraintEngineDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deploymentName := ps.ByName("deploymentname")

	log.Println("DeleteConstraintEngineDeployment", deploymentName)

	deployment, ok := ConstraintEngineDeployments[deploymentName]
	if ok {
		_, ok := ConstraintEngineProfiles[deployment.Deployment.ProfileName]
		if ok {
			undeployConstraintEngineProfile(deployment.Deployment.ProfileName)
		}

		delete(ConstraintEngineDeployments, deploymentName)

		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: ConstraintEngine deployment does not exist.")
	}
}
