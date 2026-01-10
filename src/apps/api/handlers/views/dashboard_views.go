package views

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/views/helpers"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services/filters"
	"net/http"
	"strings"
)

type DashboardViewHandlers interface {
	Dashboard(handlers.RichContext) error
}

type dashboardViewHandlers struct {
	studentService            primary.StudentPort
	internshipService         primary.InternshipPort
	timeRecordService         primary.TimeRecordPort
	internshipLocationService primary.InternshipLocationPort
}

func NewDashboardViewHandlers(
	studentService primary.StudentPort,
	internshipService primary.InternshipPort,
	timeRecordService primary.TimeRecordPort,
	internshipLocationService primary.InternshipLocationPort,
) DashboardViewHandlers {
	return &dashboardViewHandlers{
		studentService,
		internshipService,
		timeRecordService,
		internshipLocationService,
	}
}

func (h *dashboardViewHandlers) Dashboard(ctx handlers.RichContext) error {
	// Prepare filters based on user role
	studentFilters := filters.StudentFilters{}
	internshipFilters := filters.InternshipFilters{}
	timeRecordFilters := filters.TimeRecordFilters{}

	if ctx.RoleName() == "teacher" || ctx.RoleName() == "professor" {
		studentFilters.TeacherID = ctx.ProfileID()
		// internshipFilters.TeacherID = ctx.ProfileID() // Removido pois InternshipFilters não tem TeacherID
		timeRecordFilters.TeacherID = ctx.ProfileID()
	}

	// Fetch data for stats
	students, _ := h.studentService.List(studentFilters)
	internships, _ := h.internshipService.List(internshipFilters)

	// Fetch pending time records
	timeRecords, _ := h.timeRecordService.List(timeRecordFilters)

	locations, _ := h.internshipLocationService.List(filters.InternshipLocationFilters{})

	// Calculate stats
	pendingCount := 0
	recentRecords := []response.TimeRecord{}

	for i, tr := range timeRecords {
		if strings.ToUpper(tr.TimeRecordStatus().Name()) == "PENDING" {
			pendingCount++
		}
		// Get last 5 records
		if i < 5 {
			recentRecords = append(recentRecords, response.TimeRecordBuilder().BuildFromDomain(tr))
		}
	}

	data := struct {
		TotalStudents      int
		ActiveInternships  int
		PendingTimeRecords int
		TotalLocations     int
		RecentTimeRecords  []response.TimeRecord
	}{
		TotalStudents:      len(students),
		ActiveInternships:  len(internships),
		PendingTimeRecords: pendingCount,
		TotalLocations:     len(locations),
		RecentTimeRecords:  recentRecords,
	}

	return ctx.Render(http.StatusOK, "dashboard", helpers.NewPageData(ctx, "Dashboard", "dashboard", data))
}
