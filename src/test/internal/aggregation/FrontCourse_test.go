package aggregation

import (
	"course/src/domain/aggregation"
	course2 "course/src/domain/model/course"
	"course/src/domain/model/price"
	"course/src/infrastructure/GormDao"
	"course/src/test/internal/lib"
	"encoding/json"
	"testing"
)

// TestQueryDetail 测试前台课程聚合查询
func TestQueryDetail(t *testing.T) {
	courseRepo, coursePriceRepo := GormDao.NewCourseRepo(lib.DB), GormDao.NewCoursePriceRepo(lib.DB)
	course := course2.New()
	course.ID = 1
	coursePrice := price.New(price.SetCourseID(course.ID))

	fc := aggregation.NewFrontCourseAgg(course, coursePrice, courseRepo, coursePriceRepo)
	if err := fc.QueryDetail(); err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v\n%+v\n%+v", fc.Course, fc.CoursePrice, fc.CoursePrice.Price)
	}
}

// TestCreateCourse 测试创建课程
func TestCreateCourse(t *testing.T) {
	tx := lib.DB.Begin()
	courseRepo, coursePriceRepo := GormDao.NewCourseRepo(tx), GormDao.NewCoursePriceRepo(tx)
	course, coursePrice := course2.New(course2.SetName("三体")), price.New(
		price.SetMarketPrice(93),
		price.SetSalePrice(50.3),
	)

	fc := aggregation.NewFrontCourseAgg(course, coursePrice, courseRepo, coursePriceRepo)
	if err := fc.CreateCourse(); err != nil {
		tx.Rollback()
		t.Error(err)
	} else {
		tx.Commit()
		b, _ := json.Marshal(fc)
		t.Log(string(b))
	}
}
