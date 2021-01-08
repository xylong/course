package price

// Description 价格描述
type Description struct {
	Description string `json:"description" gorm:"description:'说明'"`
	IsCurrent   bool   `json:"is_current" gorm:"is_current"`
}

func NewDescription() *Description {
	return &Description{}
}

// SetDescription 设置价格说明
func SetDescription(desc string) Attr {
	return func(coursePrice *CoursePrice) {
		coursePrice.Description.Description = desc
	}
}
