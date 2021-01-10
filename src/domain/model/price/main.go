package price

import (
	"course/src/domain/model/repository"
	"time"
)

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
	Repo repository.ICoursePriceRepo `gorm:"-"`
}

// Name 设置模型名
func (cp *CoursePrice) Name() string {
	return "CoursePrice"
}

func (cp *CoursePrice) Load() error {
	return cp.Repo.FindByID(cp)
}

func New(attrs ...Attr) *CoursePrice {
	cp := &CoursePrice{
		Price:       NewPrice(),
		Description: NewDescription(),
	}
	Attrs(attrs).apply(cp)
	return cp
}

// SetCourseID 设置课程ID
func SetCourseID(id int) Attr {
	return func(coursePrice *CoursePrice) {
		coursePrice.CourseID=id
	}
}

// SetRepo 设置仓储
func SetRepo(repo repository.ICoursePriceRepo) Attr {
	return func(coursePrice *CoursePrice) {
		coursePrice.Repo=repo
	}
}