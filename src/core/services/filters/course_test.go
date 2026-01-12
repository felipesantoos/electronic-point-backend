package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCourseFilters(t *testing.T) {
	name := "Course A"
	f := CourseFilters{
		Name: &name,
	}

	assert.Equal(t, &name, f.Name)
}
