package models

type (
	// MigrationDeployParams struct
	MigrationDeployParams struct {
		SiebelMigration string `json:"SiebelMigration"`
		MigrationDesc   string `json:"MigrationDesc"`
	}

	// MigrationDeployment struct
	MigrationDeployment struct {
		Deployment            Deployment            `json:"DeploymentInfo"`
		MigrationDeployParams MigrationDeployParams `json:"MigrationDeployParams"`
	}
)
