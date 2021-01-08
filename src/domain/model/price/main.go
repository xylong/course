package price

import "time"

type (
	// 属性方法
	Attr func(*CoursePrice)
	// 属性方法集合
	Attrs []Attr
)

func (a Attrs) apply(price *CoursePrice) {
	for _, attr := range a {
		attr(price)
	}
}

// CoursePrice 课程价格
type CoursePrice struct {
	ID          int          `json:"id" gorm:"id"`
	CourseID    int          `json:"course_id" gorm:"comment:'课程id'"`
	Price       *Price       `json:"price" gorm:"embedded"`
	Description *Description `json:"description" gorm:"embedded"`
	CreatedAt   time.Time    `json:"created_at" gorm:"not null;comment:'创建时间'"`
}

func New(attrs ...Attr) *CoursePrice {
	cp := &CoursePrice{
		Price:       NewPrice(),
		Description: NewDescription(),
	}
	Attrs(attrs).apply(cp)
	return cp
}
