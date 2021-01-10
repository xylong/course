package GormDao

import (
	"course/src/domain/model/course"
	"course/src/domain/model/repository"
	"github.com/jinzhu/gorm"
)

// CourseRepo 课程仓储
type CourseRepo struct {
	db *gorm.DB
}

func NewCourseRepo(db *gorm.DB) *CourseRepo {
	return &CourseRepo{db: db}
}

// FindByID 根据id查询课程
func (r *CourseRepo) FindByID(model repository.IModel) error {
	return r.db.First(model, model.(*course.Course).ID).Error
}

// Create 添加课程
func (r *CourseRepo) Create(model repository.IModel) error {
	return r.db.Create(model).Error
}
