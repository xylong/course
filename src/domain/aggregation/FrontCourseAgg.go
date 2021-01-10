package aggregation

import (
	"course/src/domain/model/course"
	"course/src/domain/model/price"
)

// FrontCourseAgg 前台课程
type FrontCourseAgg struct {
	Course *course.Course	// 根
	CoursePrice *price.CoursePrice	// 课程价格
}
