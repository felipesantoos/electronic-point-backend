package dicontainer

import (
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services"
	"eletronic_point/src/infra/file"
	"eletronic_point/src/infra/redis"
	"eletronic_point/src/infra/repository/postgres"
)

func AccountServices() primary.AccountPort {
	repository := postgres.NewAccountRepository()
	return services.NewAccountService(repository)
}

func AuthServices() primary.AuthPort {
	repository := postgres.NewAuthPostgresRepository()
	sessionRepository := redis.NewSessionRepository()
	passwordResetRepository := redis.NewPasswordResetRepository()
	return services.NewAuthService(repository, sessionRepository, passwordResetRepository)
}

func ResourcesServices() primary.ResourcesPort {
	repository := postgres.NewResourcesRepository()
	return services.NewResourcesService(repository)
}

func StudentServices() primary.StudentPort {
	repository := postgres.NewStudentRepository()
	return services.NewStudentServices(repository)
}

func TimeRecordServices() primary.TimeRecordPort {
	repository := postgres.NewTimeRecordRepository()
	return services.NewTimeRecordServices(repository)
}

func InternshipLocationServices() primary.InternshipLocationPort {
	repository := postgres.NewInternshipLocationRepository()
	return services.NewInternshipLocationServices(repository)
}

func InternshipServices() primary.InternshipPort {
	repository := postgres.NewInternshipRepository()
	return services.NewInternshipServices(repository)
}

func TimeRecordStatusServices() primary.TimeRecordStatusPort {
	repository := postgres.NewTimeRecordStatusRepository()
	return services.NewTimeRecordStatusServices(repository)
}

func CourseServices() primary.CoursePort {
	repository := postgres.NewCourseRepository()
	return services.NewCourseServices(repository)
}

func CampusServices() primary.CampusPort {
	repository := postgres.NewCampusRepository()
	return services.NewCampusServices(repository)
}

func InstitutionServices() primary.InstitutionPort {
	repository := postgres.NewInstitutionRepository()
	return services.NewInstitutionServices(repository)
}

func FileServices() primary.FilePort {
	repository := file.NewFileRepository()
	return services.NewFileService(repository)
}
