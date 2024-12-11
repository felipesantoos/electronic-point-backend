package dicontainer

import (
	"backend_template/src/core/interfaces/usecases"
	"backend_template/src/core/services"
	"backend_template/src/infra/redis"
	"backend_template/src/infra/repository/postgres"
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
	repo := postgres.NewResourcesPostgresAdapter()
	return services.NewResourcesService(repo)
}

func StudentUseCase() usecases.StudentUseCase {
	repo := postgres.NewStudentRepository()
	return services.NewStudentService(repo)
}
