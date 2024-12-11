package dicontainer

import (
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services"
	"eletronic_point/src/infra/redis"
	"eletronic_point/src/infra/repository/postgres"
)

func AccountPort() primary.AccountPort {
	repo := postgres.NewAccountRepository()
	return services.NewAccountService(repo)
}

func AuthPort() primary.AuthPort {
	repo := postgres.NewAuthPostgresRepository()
	sessionRepo := redis.NewSessionRepository()
	passwordResetRepo := redis.NewPasswordResetRepository()
	return services.NewAuthService(repo, sessionRepo, passwordResetRepo)
}

func ResourcesPort() primary.ResourcesPort {
	repo := postgres.NewResourcesPostgresPort()
	return services.NewResourcesService(repo)
}

func StudentPort() primary.StudentPort {
	repo := postgres.NewStudentRepository()
	return services.NewStudentServices(repo)
}
