package views

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/views/helpers"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services/filters"
	"net/http"
	"strings"
	"time"
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
		timeRecordFilters.TeacherID = ctx.ProfileID()
		internshipFilters.TeacherID = ctx.ProfileID()
	}

	// Fetch data for stats
	students, _ := h.studentService.List(studentFilters)
	internships, _ := h.internshipService.List(internshipFilters)
	timeRecords, _ := h.timeRecordService.List(timeRecordFilters)
	locations, _ := h.internshipLocationService.List(filters.InternshipLocationFilters{})

	// Calculate stats and activity data
	pendingCount := 0
	recentRecords := []response.TimeRecord{}

	// Activity for the last 7 days
	activityData := make([]int, 7)
	activityLabels := make([]string, 7)
	now := time.Now()

	for i := 0; i < 7; i++ {
		day := now.AddDate(0, 0, -i)
		activityLabels[6-i] = day.Format("02/01")
	}

	for i, tr := range timeRecords {
		if strings.ToUpper(tr.TimeRecordStatus().Name()) == "PENDING" {
			pendingCount++
		}

		// Map record to activity chart
		daysAgo := int(now.Sub(tr.Date()).Hours() / 24)
		if daysAgo >= 0 && daysAgo < 7 {
			activityData[6-daysAgo]++
		}

		// Get last 5 records for the timeline
		if i < 8 {
			recentRecords = append(recentRecords, response.TimeRecordBuilder().BuildFromDomain(tr))
		}
	}

	hasActivity := false
	for _, count := range activityData {
		if count > 0 {
			hasActivity = true
			break
		}
	}

	data := struct {
		TotalStudents      int
		ActiveInternships  int
		PendingTimeRecords int
		TotalLocations     int
		RecentTimeRecords  []response.TimeRecord
		ActivityData       []int
		ActivityLabels     []string
		HasActivity        bool
	}{
		TotalStudents:      len(students),
		ActiveInternships:  len(internships),
		PendingTimeRecords: pendingCount,
		TotalLocations:     len(locations),
		RecentTimeRecords:  recentRecords,
		ActivityData:       activityData,
		ActivityLabels:     activityLabels,
		HasActivity:        hasActivity,
	}

	return ctx.Render(http.StatusOK, "dashboard", helpers.NewPageData(ctx, "Dashboard", "dashboard", data))
}
