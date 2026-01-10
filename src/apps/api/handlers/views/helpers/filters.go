package helpers

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/params"
	"eletronic_point/src/core/services/filters"
	"time"

	"github.com/google/uuid"
)

// GetStudentFilters extracts student filters from the request query parameters
func GetStudentFilters(ctx handlers.RichContext) filters.StudentFilters {
	f := filters.StudentFilters{}

	if ctx.RoleName() == "teacher" || ctx.RoleName() == "professor" {
		f.TeacherID = ctx.ProfileID()
	}

	if val := ctx.QueryParam(params.InstitutionID); val != "" {
		if id, err := uuid.Parse(val); err == nil {
			f.InstitutionID = &id
		}
	}

	if val := ctx.QueryParam(params.CampusID); val != "" {
		if id, err := uuid.Parse(val); err == nil {
			f.CampusID = &id
		}
	}

	if val := ctx.QueryParam(params.Search); val != "" {
		f.Search = &val
	}

	return f
}

// GetTimeRecordFilters extracts time record filters from the request query parameters
func GetTimeRecordFilters(ctx handlers.RichContext) filters.TimeRecordFilters {
	f := filters.TimeRecordFilters{}

	if ctx.RoleName() == "teacher" || ctx.RoleName() == "professor" {
		f.TeacherID = ctx.ProfileID()
	}

	if val := ctx.QueryParam(params.StudentID); val != "" {
		if id, err := uuid.Parse(val); err == nil {
			f.StudentID = &id
		}
	}

	if val := ctx.QueryParam(params.StatusID); val != "" {
		if id, err := uuid.Parse(val); err == nil {
			f.StatusID = &id
		}
	}

	if val := ctx.QueryParam(params.StartDate); val != "" {
		if t, err := time.Parse(time.RFC3339, val); err == nil {
			f.StartDate = &t
		} else if t, err := time.Parse("2006-01-02", val); err == nil {
			f.StartDate = &t
		}
	}

	if val := ctx.QueryParam(params.EndDate); val != "" {
		if t, err := time.Parse(time.RFC3339, val); err == nil {
			f.EndDate = &t
		} else if t, err := time.Parse("2006-01-02", val); err == nil {
			f.EndDate = &t
		}
	}

	return f
}

// GetInternshipFilters extracts internship filters from the request query parameters
func GetInternshipFilters(ctx handlers.RichContext) filters.InternshipFilters {
	f := filters.InternshipFilters{}

	if val := ctx.QueryParam(params.StudentID); val != "" {
		if id, err := uuid.Parse(val); err == nil {
			f.StudentID = &id
		}
	}

	return f
}
