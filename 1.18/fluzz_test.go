package tutorials

import (
	"testing"

	"github.com/shopspring/decimal"
)

func FuzzConvert(t *testing.F) {
	ts := []decimal.Decimal{}
	for _, tc := range ts {
		t.Add(tc)
	}
	t.Fuzz(func(tt *testing.T, ori decimal.Decimal) {
		res, err := FuzzingTest(ori.IntPart())
		if err != nil {
			tt.Error(err)
		}
		if res != "" {
			tt.Log(res)
		}

	})
}

func TestGeneratic(t *testing.T) {
	Add()
}
