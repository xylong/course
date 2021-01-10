package aggregation

import (
	"course/src/domain/model/course"
	"course/src/domain/model/price"
	"course/src/domain/model/repository"
)

// FrontCourseAgg 前台课程
type FrontCourseAgg struct {
	Course      *course.Course     // 根
	CoursePrice *price.CoursePrice // 课程价格

	CourseRepo      repository.ICourseRepo      // 课程仓储
	CoursePriceRepo repository.ICoursePriceRepo // 课程价格仓储
}

// NewFrontCourseAgg 创建前台课程聚合
func NewFrontCourseAgg(course *course.Course, coursePrice *price.CoursePrice, courseRepo repository.ICourseRepo, coursePriceRepo repository.ICoursePriceRepo) *FrontCourseAgg {
	if course == nil {
		panic("Error Root Course")
	}

	fca := &FrontCourseAgg{Course: course, CoursePrice: coursePrice, CourseRepo: courseRepo, CoursePriceRepo: coursePriceRepo}
	fca.Course.Repo, fca.CoursePrice.Repo = courseRepo, coursePriceRepo

	return fca
}

// QueryDetail 获取前台课程详情
func (c *FrontCourseAgg) QueryDetail() error {

}
