package assembler

import (
	"course/src/application/dto"
	"course/src/domain/aggregation"
)

type CourseResponse struct {
}

// M2D_SimpleCourseInfo 模型转dto
func (req *CourseResponse) M2D_SimpleCourseInfo(front *aggregation.FrontCourse) *dto.SimpleCourseRep {
	c, p := front.Course, front.CoursePrice
	simpleCourse := &dto.SimpleCourseRep{}
	simpleCourse.ID = c.ID
	simpleCourse.Name = c.Info.Name
	simpleCourse.MarketPrice = p.Price.MarketPrice
	simpleCourse.SalePrice = p.Price.SalePrice
	return simpleCourse
}
