package helpers

import (
	eletronic_point_campus "eletronic_point/src/core/domain/campus"
	eletronic_point_course "eletronic_point/src/core/domain/course"
	eletronic_point_institution "eletronic_point/src/core/domain/institution"
	eletronic_point_internship_location "eletronic_point/src/core/domain/internshipLocation"
	eletronic_point_role "eletronic_point/src/core/domain/role"
	eletronic_point_student "eletronic_point/src/core/domain/student"
	eletronic_point_time_record_status "eletronic_point/src/core/domain/timeRecordStatus"
)

type Option struct {
	Label string
	Value string
}

func ToOptions(list interface{}) []Option {
	options := make([]Option, 0)

	if list == nil {
		return options
	}

	switch v := list.(type) {
	case []eletronic_point_campus.Campus:
		for _, item := range v {
			options = append(options, Option{Label: item.Name(), Value: item.ID().String()})
		}
	case []eletronic_point_course.Course:
		for _, item := range v {
			options = append(options, Option{Label: item.Name(), Value: item.ID().String()})
		}
	case []eletronic_point_institution.Institution:
		for _, item := range v {
			options = append(options, Option{Label: item.Name(), Value: item.ID().String()})
		}
	case []eletronic_point_student.Student:
		for _, item := range v {
			id := ""
			if item.ID() != nil {
				id = item.ID().String()
			}
			options = append(options, Option{Label: item.Name(), Value: id})
		}
	case []eletronic_point_internship_location.InternshipLocation:
		for _, item := range v {
			options = append(options, Option{Label: item.Name(), Value: item.ID().String()})
		}
	case []eletronic_point_time_record_status.TimeRecordStatus:
		for _, item := range v {
			options = append(options, Option{Label: item.Name(), Value: item.ID().String()})
		}
	case []eletronic_point_role.Role:
		for _, item := range v {
			options = append(options, Option{Label: item.Name(), Value: item.ID().String()})
		}
	}
	return options
}
