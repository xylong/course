package repository

// ICourseRepo 课程仓储
type ICourseRepo interface {
	// 根据id查找
	FindByID(IModel) error
}
