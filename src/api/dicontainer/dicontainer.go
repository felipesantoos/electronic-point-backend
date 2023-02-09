package dicontainer

import (
	"dit_backend/src/core/interfaces/usecases"
	"dit_backend/src/core/services"
	"dit_backend/src/infra/repository/postgres"
	"dit_backend/src/infra/repository/redis"
	"os"
)

func AccountUseCase() usecases.AccountUseCase {
	repo := postgres.NewAccountRepository()
	return services.NewAccountService(repo)
}

func AuthUseCase() usecases.AuthUseCase {
	repo := postgres.NewAuthPostgres()
	sessionRepo := redis.NewSessionRepository()
	passwordResetRepo := redis.NewPasswordResetRepository()
	return services.NewAuthService(repo, sessionRepo, passwordResetRepo)
}

func ResourcesUseCase() usecases.ResourcesUseCase {
	repo := postgres.NewResourcesPostgresAdapter()
	return services.NewResourcesService(repo)
}

func getAppType() string {
	return os.Getenv("APPLICATION_TYPE")
}
