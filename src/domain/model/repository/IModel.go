package repository

// IModel 实体
type IModel interface {
	// 设置实体名
	Name() string
	// 载入数据
	Load() error
}
