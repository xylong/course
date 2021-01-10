package repository

// ICoursePrice 课程价格操作
type ICoursePriceRepo interface {
	// 根据id查找
	FindByID(IModel) error
}
