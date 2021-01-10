package aggregation

import (
	"course/src/domain/model/course"
	"course/src/domain/model/price"
	"course/src/domain/model/repository"
	"course/src/infrastructure/errors"
)

// FrontCourse 前台课程
type FrontCourse struct {
	Course      *course.Course     // 根
	CoursePrice *price.CoursePrice // 课程价格

	CourseRepo      repository.ICourseRepo      // 课程仓储
	CoursePriceRepo repository.ICoursePriceRepo // 课程价格仓储
}

// NewFrontCourseAgg 创建前台课程聚合
func NewFrontCourseAgg(course *course.Course, coursePrice *price.CoursePrice, courseRepo repository.ICourseRepo, coursePriceRepo repository.ICoursePriceRepo) *FrontCourse {
	if course == nil {
		panic("Error Root Course")
	}

	fc := &FrontCourse{Course: course, CoursePrice: coursePrice, CourseRepo: courseRepo, CoursePriceRepo: coursePriceRepo}
	fc.Course.Repo, fc.CoursePrice.Repo = courseRepo, coursePriceRepo

	return fc
}

// QueryDetail 获取前台课程详情
func (fc *FrontCourse) QueryDetail() error {
	if fc.Course.ID <= 0 {
		return errors.NewNotFoundIDError("Course")
	}
	// 课程信息
	if err := fc.Course.Load(); err != nil {
		return errors.NewNotFoundDataError("Course", err.Error())
	}
	// 价格信息
	if err := fc.CoursePrice.Load(); err != nil {
		return errors.NewNotFoundDataError("CoursePrise", err.Error())
	}

	return nil
}

// CreateCourse 创建课程
func (fc *FrontCourse) CreateCourse() error {
	if err := fc.Course.Create(); err != nil {
		return errors.NewCreateError("Course", err.Error())
	}
	fc.CoursePrice.CourseID = fc.Course.ID
	if err := fc.CoursePrice.Create(); err != nil {
		return errors.NewCreateError("CoursePrice", err.Error())
	}

	return nil
}
