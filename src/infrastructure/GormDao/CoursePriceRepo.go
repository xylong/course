package GormDao

import (
	"course/src/domain/model/price"
	"course/src/domain/model/repository"
	"github.com/jinzhu/gorm"
)

// CoursePriceRepo 课程价格实体仓储
type CoursePriceRepo struct {
	db *gorm.DB
}

func NewCoursePriceRepo(db *gorm.DB) *CoursePriceRepo {
	return &CoursePriceRepo{db: db}
}

// FindByID 根据id查找课程价格
func (r *CoursePriceRepo) FindByID(model repository.IModel) error {
	return r.db.First(model, model.(*price.CoursePrice).CourseID).Error
}

// Create 添加课程价格信息
func (r *CoursePriceRepo) Create(model repository.IModel) error {
	return r.db.Create(model).Error
}
