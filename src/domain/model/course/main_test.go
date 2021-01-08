package course

import (
	"encoding/json"
	"fmt"
	"testing"
)

// TestNew 测试创建课程
func TestNew(t *testing.T) {
	c := New(
		SetName("apple"),
		SetImage("banana"),
		SetTime(3679),
	)
	b, _ := json.Marshal(c)
	fmt.Println(string(b))

	c.Time.Add(100)
	_b, _ := json.Marshal(c)
	fmt.Println(string(_b))
}
