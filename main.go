package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"./restapis"
	"github.com/julienschmidt/httprouter"
)

var port = ":8889"

// Index - For /
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	// CG Bootstrap
	router.GET("/siebel/v1.0/cginfo", restapis.GetCGInfo)
	router.POST("/siebel/v1.0/cginfo", restapis.PostCGInfo)
	router.GET("/siebel/v1.0/cloudgateway/bootstrapCG", restapis.GetBootstrapCG)
	router.POST("/siebel/v1.0/cloudgateway/bootstrapCG", restapis.PostBootstrapCG)

	// List APIs
	router.GET("/siebel/v1.0/appicon", restapis.GetAppIcon)
	router.GET("/siebel/v1.0/cloudgateway/discovery/services", restapis.GetDiscoveryServices)

	// Security Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/security", restapis.PostSecurityProfile)
	router.GET("/siebel/v1.0/cloudgateway/profiles/security", restapis.GetSecurityProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/security/:name", restapis.GetSecurityProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/security/:name", restapis.PutSecurityProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/security/:name", restapis.DeleteSecurityProfile)

	// Enterprise Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/enterprises", restapis.PostEnterpriseProfile)
	router.GET("/siebel/v1.0/cloudgateway/profiles/enterprises", restapis.GetEnterpriseProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/enterprises/:name", restapis.GetEnterpriseProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/enterprises/:name", restapis.PutEnterpriseProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/enterprises/:name", restapis.DeleteEnterpriseProfile)

	// Enterprise Deployment APIs
	router.POST("/siebel/v1.0/cloudgateway/deployments/enterprises", restapis.PostEnterpriseDeployment)
	router.GET("/siebel/v1.0/cloudgateway/deployments/enterprises", restapis.GetEnterpriseDeployments)
	router.GET("/siebel/v1.0/cloudgateway/deployments/enterprises/:name", restapis.GetEnterpriseDeployment)
	router.PUT("/siebel/v1.0/cloudgateway/deployments/enterprises/:name", restapis.PutEnterpriseDeployment)
	router.DELETE("/siebel/v1.0/cloudgateway/deployments/enterprises/:name", restapis.DeleteEnterpriseDeployment)

	// Server Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/servers", restapis.PostServerProfile)
	router.GET("/siebel/v1.0/cloudgateway/profiles/servers", restapis.GetServerProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/servers/:name", restapis.GetServerProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/servers/:name", restapis.PutServerProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/servers/:name", restapis.DeleteServerProfile)

	// Server Deployment APIs
	router.POST("/siebel/v1.0/cloudgateway/deployments/servers", restapis.PostServerDeployment)
	router.GET("/siebel/v1.0/cloudgateway/deployments/servers", restapis.GetServerDeployments)
	router.GET("/siebel/v1.0/cloudgateway/deployments/servers/:name", restapis.GetServerDeployment)
	router.PUT("/siebel/v1.0/cloudgateway/deployments/servers/:name", restapis.PutServerDeployment)
	router.DELETE("/siebel/v1.0/cloudgateway/deployments/servers/:name", restapis.DeleteServerDeployment)

	// SWSM Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/swsm", restapis.PostSWSMProfile)
	router.GET("/siebel/v1.0/cloudgateway/profiles/swsm", restapis.GetSWSMProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/swsm/:name", restapis.GetSWSMProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/swsm/:name", restapis.PutSWSMProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/swsm/:name", restapis.DeleteSWSMProfile)

	// SWSM Deployment APIs
	router.POST("/siebel/v1.0/cloudgateway/deployments/swsm", restapis.PostSWSMDeployment)
	router.GET("/siebel/v1.0/cloudgateway/deployments/swsm", restapis.GetSWSMDeployments)
	router.GET("/siebel/v1.0/cloudgateway/deployments/swsm/:name", restapis.GetSWSMDeployment)
	router.PUT("/siebel/v1.0/cloudgateway/deployments/swsm/:name", restapis.PutSWSMDeployment)
	router.DELETE("/siebel/v1.0/cloudgateway/deployments/swsm/:name", restapis.DeleteSWSMDeployment)

	// CacheServer Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/cacheserver", restapis.PostCacheServerProfile)
	router.POST("/siebel/v1.0/cloudgateway/profiles/cacheserver/tangoconfig", restapis.PostCacheServerTangoConfigConversion)
	router.POST("/siebel/v1.0/cloudgateway/profiles/cacheserver/cacheconfig", restapis.PostCacheServerCacheConfigConversion)
	router.GET("/siebel/v1.0/cloudgateway/profiles/cacheserver", restapis.GetCacheServerProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/cacheserver/:name", restapis.GetCacheServerProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/cacheserver/:name", restapis.PutCacheServerProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/cacheserver/:name", restapis.DeleteCacheServerProfile)

	// CacheServer Deployment APIs
	router.POST("/siebel/v1.0/cloudgateway/deployments/cacheserver", restapis.PostCacheServerDeployment)
	router.GET("/siebel/v1.0/cloudgateway/deployments/cacheserver", restapis.GetCacheServerDeployments)
	router.GET("/siebel/v1.0/cloudgateway/deployments/cacheserver/:name", restapis.GetCacheServerDeployment)
	router.PUT("/siebel/v1.0/cloudgateway/deployments/cacheserver/:name", restapis.PutCacheServerDeployment)
	router.DELETE("/siebel/v1.0/cloudgateway/deployments/cacheserver/:name", restapis.DeleteCacheServerDeployment)

	// CacheClient Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/cacheclient", restapis.PostCacheClientProfile)
	router.POST("/siebel/v1.0/cloudgateway/profiles/cacheclient/cacheconfig", restapis.PostCacheClientCacheConfigConversion)
	router.GET("/siebel/v1.0/cloudgateway/profiles/cacheclient", restapis.GetCacheClientProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/cacheclient/:name", restapis.GetCacheClientProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/cacheclient/:name", restapis.PutCacheClientProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/cacheclient/:name", restapis.DeleteCacheClientProfile)

	// Migration Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/migrations", restapis.PostMigrationProfile)
	router.GET("/siebel/v1.0/cloudgateway/profiles/migrations", restapis.GetMigrationProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/migrations/:name", restapis.GetMigrationProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/migrations/:name", restapis.PutMigrationProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/migrations/:name", restapis.DeleteMigrationProfile)

	// Migration Deployment APIs
	router.POST("/siebel/v1.0/cloudgateway/deployments/migrations", restapis.PostMigrationDeployment)
	router.GET("/siebel/v1.0/cloudgateway/deployments/migrations", restapis.GetMigrationDeployments)
	router.GET("/siebel/v1.0/cloudgateway/deployments/migrations/:name", restapis.GetMigrationDeployment)
	router.PUT("/siebel/v1.0/cloudgateway/deployments/migrations/:name", restapis.PutMigrationDeployment)
	router.DELETE("/siebel/v1.0/cloudgateway/deployments/migrations/:name", restapis.DeleteMigrationDeployment)

	// Constraint Engine Profile APIs
	router.POST("/siebel/v1.0/cloudgateway/profiles/constraintengine", restapis.PostConstraintEngineProfile)
	router.GET("/siebel/v1.0/cloudgateway/profiles/constraintengine", restapis.GetConstraintEngineProfiles)
	router.GET("/siebel/v1.0/cloudgateway/profiles/constraintengine/:name", restapis.GetConstraintEngineProfile)
	router.PUT("/siebel/v1.0/cloudgateway/profiles/constraintengine/:name", restapis.PutConstraintEngineProfile)
	router.DELETE("/siebel/v1.0/cloudgateway/profiles/constraintengine/:name", restapis.DeleteConstraintEngineProfile)

	// Constraint Engine Deployment APIs
	router.POST("/siebel/v1.0/cloudgateway/deployments/constraintengine", restapis.PostConstraintEngineDeployment)
	router.GET("/siebel/v1.0/cloudgateway/deployments/constraintengine", restapis.GetConstraintEngineDeployments)
	router.GET("/siebel/v1.0/cloudgateway/deployments/constraintengine/:name", restapis.GetConstraintEngineDeployment)
	router.PUT("/siebel/v1.0/cloudgateway/deployments/constraintengine/:name", restapis.PutConstraintEngineDeployment)
	router.DELETE("/siebel/v1.0/cloudgateway/deployments/constraintengine/:name", restapis.DeleteConstraintEngineDeployment)

	// Cache Management APIs
	router.GET("/siebel/v1.0/cloudgateway/cache", restapis.GetCaches)
	router.GET("/siebel/v1.0/cloudgateway/cache/:cachename", restapis.GetCache)
	router.DELETE("/siebel/v1.0/cloudgateway/cache", restapis.DeleteCaches)
	router.DELETE("/siebel/v1.0/cloudgateway/cache/:cachename", restapis.DeleteCache)
	router.DELETE("/siebel/v1.0/cloudgateway/clearcache", restapis.ClearCaches)
	router.DELETE("/siebel/v1.0/cloudgateway/clearcache/:cachename", restapis.ClearCache)

	// Seed Cache Generator
	rand.Seed(time.Now().UTC().UnixNano())

	// Loading data
	restapis.LoadAll()

	// Start server
	log.Println("Starting server...")
	log.Println("http://localhost" + port + "/siebel/v1.0/")
	log.Fatal(http.ListenAndServe(port, router))
}
