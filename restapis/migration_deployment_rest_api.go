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

// MigrationDeployments map
var MigrationDeployments = make(map[string]models.MigrationDeployment)

func deployMigrationProfile(profileName string) {
	profile, ok := MigrationProfiles[profileName]
	if ok {
		profile.Profile.Deploy()
		MigrationProfiles[profileName] = profile
	}
}

func undeployMigrationProfile(profileName string) {
	profile, ok := MigrationProfiles[profileName]
	if ok {
		profile.Profile.Undeploy()
		MigrationProfiles[profileName] = profile
	}
}

// PostMigrationDeployment - POST /deployments/migrations
func PostMigrationDeployment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deployment := models.MigrationDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	logger.Logger("PostMigrationDeployment", string(msg))

	_, ok := MigrationProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Inexist Deployment Profile.")
		return
	}

	deploymentName := deployment.MigrationDeployParams.SiebelMigration
	if len(deploymentName) > 0 {
		_, ok := MigrationDeployments[deploymentName]
		if ok {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Error: Enterprise deployment with same name already exists.")
		} else {
			deployment.Deployment.Check()
			MigrationDeployments[deploymentName] = deployment

			deployMigrationProfile(deployment.Deployment.ProfileName)

			w.WriteHeader(201)
			fmt.Fprintf(w, "Succeed.")
		}
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Name.")
	}
}

// GetMigrationDeployments - GET /deployments/migrations
func GetMigrationDeployments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var deploymentsJSON bytes.Buffer
	deploymentsJSON.WriteString("{\"MigrationDeployment\":[")

	logger.Logger("GetMigrationDeployments", strconv.Itoa(len(MigrationDeployments)))

	ok := false
	for _, deployment := range MigrationDeployments {
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

// GetMigrationDeployment - GET /deployments/migrations/:deploymentname
func GetMigrationDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment, ok := MigrationDeployments[ps.ByName("deploymentname")]

	logger.Logger("GetMigrationDeployment", ps.ByName("deploymentname"))

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

// PutMigrationDeployment - PUT /deployments/migrations/:deploymentname
func PutMigrationDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deployment := models.MigrationDeployment{}
	json.NewDecoder(r.Body).Decode(&deployment)

	msg, _ := json.Marshal(deployment)
	logger.Logger("PutMigrationDeployment", string(msg))

	_, ok := MigrationProfiles[deployment.Deployment.ProfileName]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: Expecting Deployment Profile.")
		return
	}

	deploymentName := ps.ByName("deploymentname")
	_, ok = MigrationDeployments[deploymentName]
	if ok {
		if deployment.MigrationDeployParams.SiebelMigration == deploymentName {
			oldProfileName := MigrationDeployments[deploymentName].Deployment.ProfileName
			if deployment.Deployment.ProfileName != oldProfileName {
				deployMigrationProfile(deployment.Deployment.ProfileName)
				undeployMigrationProfile(oldProfileName)
			}

			MigrationDeployments[deploymentName] = deployment

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

// DeleteMigrationDeployment - DELETE /deployments/migrations/:deploymentname
func DeleteMigrationDeployment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deploymentName := ps.ByName("deploymentname")

	logger.Logger("DeleteMigrationDeployment", deploymentName)

	deployment, ok := MigrationDeployments[deploymentName]
	if ok {
		_, ok := MigrationProfiles[deployment.Deployment.ProfileName]
		if ok {
			undeployMigrationProfile(deployment.Deployment.ProfileName)
		}

		delete(MigrationDeployments, deploymentName)

		w.WriteHeader(200)
		fmt.Fprintf(w, "Succeed.")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Enterprise deployment does not exist.")
	}
}
