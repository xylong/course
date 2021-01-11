package service

import (
	"course/src/application/assembler"
	"course/src/application/dto"
	"course/src/domain/aggregation"
	"course/src/domain/model/price"
	"course/src/infrastructure/GormDao"
	"github.com/jinzhu/gorm"
)

// CourseService 课程业务
type CourseService struct {
	Req *assembler.CourseRequest
	Rep *assembler.CourseResponse
	DB  *gorm.DB
}

// GetSimpleCourse 获取简洁课程信息
func (s *CourseService) GetSimpleCourse(req *dto.SimpleCourseReq) *dto.SimpleCourseRep {
	c := s.Req.D2M_Course(req)
	cp := price.New(price.SetCourseID(c.ID))
	courseRepo, coursePriceRepo := GormDao.NewCourseRepo(s.DB), GormDao.NewCoursePriceRepo(s.DB)
	frontCourse := aggregation.NewFrontCourseAgg(c, cp, courseRepo, coursePriceRepo)
	err := frontCourse.QueryDetail()
	if err != nil {
		panic(err)
	}

	return s.Rep.M2D_SimpleCourseInfo(frontCourse)
}
