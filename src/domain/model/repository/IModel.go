package repository

// IModel 实体接口
type IModel interface {
	// 设置实体名
	Name() string
	// 载入数据
	Load() error
	// 添加课程
	Create() error
}
