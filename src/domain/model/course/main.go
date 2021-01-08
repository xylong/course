package course

type (
	// Attr 属性方法
	Attr func(*Course)
	// 属性集合
	Attrs []Attr
)

// apply 设置属性
func (a Attrs) apply(course *Course) {
	for _, attr := range a {
		attr(course)
	}
}

// Course 课程
type Course struct {
	ID        int        `json:"id" gorm:"id"`
	Info      *Info      `json:"info" gorm:"embedded"`       // 课程信息
	Time      *Time      `json:"time" gorm:"embedded"`       // 时长
	CreatedAt *CreatedAt `json:"created_at" gorm:"embedded"` // 创建时间
}

// New 创建课程
func New(attrs ...Attr) *Course {
	course := &Course{
		Info:      NewInfo(),
		Time:      NewTime(),
		CreatedAt: NewCreatedAt(),
	}
	Attrs(attrs).apply(course)
	return course
}
