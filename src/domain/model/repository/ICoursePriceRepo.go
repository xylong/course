package repository

// ICoursePrice 课程价格仓储接口
type ICoursePriceRepo interface {
	// 根据id查找
	FindByID(IModel) error
}
