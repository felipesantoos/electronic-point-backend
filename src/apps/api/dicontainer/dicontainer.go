package dicontainer

import (
	"eletronic_point/src/core/interfaces/usecases"
	"eletronic_point/src/core/services"
	"eletronic_point/src/infra/redis"
	"eletronic_point/src/infra/repository/postgres"
)

func AccountUseCase() usecases.AccountUseCase {
	repo := postgres.NewAccountRepository()
	return services.NewAccountService(repo)
}

func AuthUseCase() usecases.AuthUseCase {
	repo := postgres.NewAuthPostgresRepository()
	sessionRepo := redis.NewSessionRepository()
	passwordResetRepo := redis.NewPasswordResetRepository()
	return services.NewAuthService(repo, sessionRepo, passwordResetRepo)
}

func ResourcesUseCase() usecases.ResourcesUseCase {
	repo := postgres.NewResourcesPostgresPort()
	return services.NewResourcesService(repo)
}

func StudentUseCase() usecases.StudentUseCase {
	repo := postgres.NewStudentRepository()
	return services.NewStudentService(repo)
}
