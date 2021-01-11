package assembler

import (
	"course/src/application/dto"
	"course/src/domain/model/course"
)

// CourseRequest 课程请求
type CourseRequest struct {
}

// D2M_Course dto转模型
func (req *CourseRequest) D2M_Course(courseReq *dto.SimpleCourseReq) *course.Course {
	c := course.New()
	c.ID = courseReq.ID
	return c
}
