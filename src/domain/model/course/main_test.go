package course

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	c := New(
		SetName("apple"),
		SetImage("banana"),
		SetTime(12345),
	)
	b, _ := json.Marshal(c)
	fmt.Println(string(b))

	c.Time.Add(100)
	_b, _ := json.Marshal(c)
	fmt.Println(string(_b))
}
