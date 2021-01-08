package price

import (
	"encoding/json"
	"testing"
)

func TestNew(t *testing.T) {
	cp := New(
		SetMarketPrice(25),
		SetSalePrice(20),
		SetDescription("测试"),
	)
	if b, err := json.Marshal(cp); err != nil {
		t.Error(err)
	} else {
		t.Log(string(b))
	}
}
