package repository

// ICourseRepo 课程操作
type ICourseRepo interface {
	// 根据id查找
	FindByID(IModel) error
}
