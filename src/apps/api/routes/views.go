package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers/views"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type viewRouter struct {
	authHandlers               views.AuthViewHandlers
	dashboardHandlers          views.DashboardViewHandlers
	resourceHandlers           views.ResourceViewHandlers
	accountHandlers            views.AccountViewHandlers
	studentHandlers            views.StudentViewHandlers
	internshipLocationHandlers views.InternshipLocationViewHandlers
	internshipHandlers         views.InternshipViewHandlers
	timeRecordHandlers         views.TimeRecordViewHandlers
	timeRecordStatusHandlers   views.TimeRecordStatusViewHandlers
}

func NewViewRouter() Router {
	authService := dicontainer.AuthServices()
	studentService := dicontainer.StudentServices()
	internshipService := dicontainer.InternshipServices()
	timeRecordService := dicontainer.TimeRecordServices()
	internshipLocationService := dicontainer.InternshipLocationServices()
	institutionService := dicontainer.InstitutionServices()
	campusService := dicontainer.CampusServices()
	courseService := dicontainer.CourseServices()
	accountService := dicontainer.AccountServices()
	resourcesService := dicontainer.ResourcesServices()
	timeRecordStatusService := dicontainer.TimeRecordStatusServices()

	return &viewRouter{
		authHandlers: views.NewAuthViewHandlers(authService),
		dashboardHandlers: views.NewDashboardViewHandlers(
			studentService,
			internshipService,
			timeRecordService,
			internshipLocationService,
		),
		resourceHandlers: views.NewResourceViewHandlers(
			institutionService,
			campusService,
			courseService,
		),
		accountHandlers:            views.NewAccountViewHandlers(accountService, resourcesService),
		studentHandlers:            views.NewStudentViewHandlers(studentService, institutionService, campusService, courseService, accountService, resourcesService),
		internshipLocationHandlers: views.NewInternshipLocationViewHandlers(internshipLocationService, internshipService),
		internshipHandlers:         views.NewInternshipViewHandlers(internshipService, studentService, internshipLocationService, accountService, resourcesService),
		timeRecordHandlers:         views.NewTimeRecordViewHandlers(timeRecordService, studentService, timeRecordStatusService, accountService, resourcesService),
		timeRecordStatusHandlers:   views.NewTimeRecordStatusViewHandlers(timeRecordStatusService),
	}
}

func (r *viewRouter) Load(rootEndpoint *echo.Group) {
	// Public routes
	rootEndpoint.GET("/login", middlewares.EnhanceContext(r.authHandlers.LoginPage))
	rootEndpoint.POST("/login", middlewares.EnhanceContext(r.authHandlers.Login))
	rootEndpoint.POST("/logout", middlewares.EnhanceContext(r.authHandlers.Logout))
	rootEndpoint.GET("/reset-password", middlewares.EnhanceContext(r.authHandlers.ResetPasswordPage))
	rootEndpoint.POST("/reset-password", middlewares.EnhanceContext(r.authHandlers.AskPasswordResetMail))
	rootEndpoint.GET("/reset-password/:token", middlewares.EnhanceContext(r.authHandlers.ResetPasswordConfirmPage))
	rootEndpoint.PUT("/reset-password/:token", middlewares.EnhanceContext(r.authHandlers.UpdatePasswordByPasswordReset))

	// Protected routes
	protected := rootEndpoint.Group("", middlewares.ViewsAuthorize)

	protected.GET("/", middlewares.EnhanceContext(r.dashboardHandlers.Dashboard))

	// Profile
	protected.GET("/accounts/profile", middlewares.EnhanceContext(r.accountHandlers.ProfilePage))
	protected.PUT("/accounts/profile", middlewares.EnhanceContext(r.accountHandlers.UpdateProfile))
	protected.PUT("/accounts/update-password", middlewares.EnhanceContext(r.accountHandlers.UpdatePassword))

	// Admin Accounts
	admin := protected.Group("/admin", middlewares.AdminAuthorize)
	admin.GET("/accounts", middlewares.EnhanceContext(r.accountHandlers.List))
	admin.GET("/accounts/new", middlewares.EnhanceContext(r.accountHandlers.CreatePage))
	admin.POST("/accounts", middlewares.EnhanceContext(r.accountHandlers.Create))
	admin.GET("/accounts/:id", middlewares.EnhanceContext(r.accountHandlers.Show))
	admin.GET("/accounts/:id/edit", middlewares.EnhanceContext(r.accountHandlers.EditPage))
	admin.PUT("/accounts/:id", middlewares.EnhanceContext(r.accountHandlers.Update))

	// Students
	protected.GET("/students", middlewares.EnhanceContext(r.studentHandlers.List))
	protected.GET("/students/new", middlewares.EnhanceContext(r.studentHandlers.CreatePage))
	protected.POST("/students", middlewares.EnhanceContext(r.studentHandlers.Create))
	protected.GET("/students/:id", middlewares.EnhanceContext(r.studentHandlers.Show))
	protected.GET("/students/:id/edit", middlewares.EnhanceContext(r.studentHandlers.EditPage))
	protected.PUT("/students/:id", middlewares.EnhanceContext(r.studentHandlers.Update))

	// Internship Locations
	protected.GET("/internship-locations", middlewares.EnhanceContext(r.internshipLocationHandlers.List))
	protected.GET("/internship-locations/new", middlewares.EnhanceContext(r.internshipLocationHandlers.CreatePage))
	protected.POST("/internship-locations", middlewares.EnhanceContext(r.internshipLocationHandlers.Create))
	protected.GET("/internship-locations/:id", middlewares.EnhanceContext(r.internshipLocationHandlers.Show))
	protected.GET("/internship-locations/:id/edit", middlewares.EnhanceContext(r.internshipLocationHandlers.EditPage))
	protected.PUT("/internship-locations/:id", middlewares.EnhanceContext(r.internshipLocationHandlers.Update))

	// Internships
	protected.GET("/internships", middlewares.EnhanceContext(r.internshipHandlers.List))
	protected.GET("/internships/new", middlewares.EnhanceContext(r.internshipHandlers.CreatePage))
	protected.POST("/internships", middlewares.EnhanceContext(r.internshipHandlers.Create))
	protected.GET("/internships/:id", middlewares.EnhanceContext(r.internshipHandlers.Show))

	// Time Records
	protected.GET("/time-records", middlewares.EnhanceContext(r.timeRecordHandlers.List))
	protected.GET("/time-records/new", middlewares.EnhanceContext(r.timeRecordHandlers.CreatePage))
	protected.POST("/time-records", middlewares.EnhanceContext(r.timeRecordHandlers.Create))
	protected.GET("/time-records/:id", middlewares.EnhanceContext(r.timeRecordHandlers.Show))
	protected.PATCH("/time-records/:id/approve", middlewares.EnhanceContext(r.timeRecordHandlers.Approve))
	protected.PATCH("/time-records/:id/disapprove", middlewares.EnhanceContext(r.timeRecordHandlers.Disapprove))

	// Time Record Status
	protected.GET("/time-record-status", middlewares.EnhanceContext(r.timeRecordStatusHandlers.List))
	protected.GET("/time-record-status/new", middlewares.EnhanceContext(r.timeRecordStatusHandlers.CreatePage), middlewares.AdminAuthorize)
	protected.POST("/time-record-status", middlewares.EnhanceContext(r.timeRecordStatusHandlers.Create), middlewares.AdminAuthorize)
	protected.GET("/time-record-status/:id", middlewares.EnhanceContext(r.timeRecordStatusHandlers.Show))
	protected.GET("/time-record-status/:id/edit", middlewares.EnhanceContext(r.timeRecordStatusHandlers.EditPage), middlewares.AdminAuthorize)
	protected.PUT("/time-record-status/:id", middlewares.EnhanceContext(r.timeRecordStatusHandlers.Update), middlewares.AdminAuthorize)

	// Resources (Institutions, Campus, Courses)
	protected.GET("/institutions", middlewares.EnhanceContext(r.resourceHandlers.Institutions))
	protected.GET("/institutions/new", middlewares.EnhanceContext(r.resourceHandlers.CreateInstitutionPage), middlewares.AdminAuthorize)
	protected.POST("/institutions", middlewares.EnhanceContext(r.resourceHandlers.CreateInstitution), middlewares.AdminAuthorize)
	protected.GET("/institutions/:id/edit", middlewares.EnhanceContext(r.resourceHandlers.EditInstitutionPage), middlewares.AdminAuthorize)
	protected.PUT("/institutions/:id", middlewares.EnhanceContext(r.resourceHandlers.UpdateInstitution), middlewares.AdminAuthorize)

	protected.GET("/campus", middlewares.EnhanceContext(r.resourceHandlers.Campus))
	protected.GET("/campus/new", middlewares.EnhanceContext(r.resourceHandlers.CreateCampusPage), middlewares.AdminAuthorize)
	protected.POST("/campus", middlewares.EnhanceContext(r.resourceHandlers.CreateCampus), middlewares.AdminAuthorize)
	protected.GET("/campus/:id/edit", middlewares.EnhanceContext(r.resourceHandlers.EditCampusPage), middlewares.AdminAuthorize)
	protected.PUT("/campus/:id", middlewares.EnhanceContext(r.resourceHandlers.UpdateCampus), middlewares.AdminAuthorize)

	protected.GET("/courses", middlewares.EnhanceContext(r.resourceHandlers.Courses))
	protected.GET("/courses/new", middlewares.EnhanceContext(r.resourceHandlers.CreateCoursePage), middlewares.AdminAuthorize)
	protected.POST("/courses", middlewares.EnhanceContext(r.resourceHandlers.CreateCourse), middlewares.AdminAuthorize)
	protected.GET("/courses/:id/edit", middlewares.EnhanceContext(r.resourceHandlers.EditCoursePage), middlewares.AdminAuthorize)
	protected.PUT("/courses/:id", middlewares.EnhanceContext(r.resourceHandlers.UpdateCourse), middlewares.AdminAuthorize)
}
