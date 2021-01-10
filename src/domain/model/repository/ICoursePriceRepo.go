package repository

// ICoursePrice 课程价格仓储接口
type ICoursePriceRepo interface {
	// 根据id查找
	FindByID(IModel) error
	// 创建课程价格信息
	Create(model IModel) error
}
