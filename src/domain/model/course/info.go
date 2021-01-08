package course

// Info 课程信息
type Info struct {
	Name  string `json:"name" gorm:"type:varchar(100);not null;comment:'名称'"`
	Image string `json:"image" gorm:"type:varchar(100);comment:'图片'"`
}

func NewInfo() *Info {
	return &Info{}
}

// SetName 设置名称
func SetName(name string) Attr {
	return func(course *Course) {
		if name != "" {
			course.Info.Name = name
		}
	}
}

// SetImage 设置图片
func SetImage(image string) Attr {
	return func(course *Course) {
		if image != "" {
			course.Info.Image = image
		}
	}
}
