package course

import (
	"encoding/json"
	"time"
)

// CreatedAt 创建时间
type CreatedAt struct {
	Date time.Time `json:"created_at" gorm:"column:created_at;not null;comment:'创建时间'"`
}

// MarshalJSON 实现json接口
func (c CreatedAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Date    time.Time `json:"date"`
		Default string    `json:"default"`
		Unix    int64     `json:"unix"`
	}{
		Date:    c.Date,
		Default: c.Format(),
		Unix:    c.Date.Unix(),
	})
}

// Format 格式化时间
func (c *CreatedAt) Format() string {
	return c.Date.Format("2006-01-02 15:04:05")
}

func NewCreatedAt() *CreatedAt {
	return &CreatedAt{
		Date: time.Now(),
	}
}
