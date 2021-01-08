package course

import (
	"encoding/json"
	"fmt"
)

// Time 课程时长
type Time struct {
	Second int64 `json:"seconds" gorm:"type:bigint(20);default:0;comment:'时长'"`
}

func NewTime() *Time {
	return &Time{}
}

// Add 增加时长
func (t *Time) Add(sec int64) {
	t.Second += sec
}

// MarshalJSON 实现json接口
func (t *Time) MarshalJSON() ([]byte, error) {
	seconds := t.Second
	hour := seconds / 3600
	minute := (seconds - hour*3600) / 60
	second := seconds - hour*3600 - minute*60

	return json.Marshal(map[string]interface{}{
		"Second": seconds,
		"Format": fmt.Sprintf("%d时%d分%d秒", hour, minute, second),
	})
}

// SetTime 设置时长
func SetTime(sec int64) Attr {
	return func(course *Course) {
		if sec != 0 {
			course.Time.Second = sec
		}
	}
}
