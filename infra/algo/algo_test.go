package algo

import (
	"testing"

	. "github.com/smgocartystreets/goconvey"
)

func TestRandomValue(t *testing.T) {
	count, amount := 10, 10000
	remainCount, remainAmount := count, amount
	sumAmount := 0
	for i := 0; i < count; i++ {
		curAmount := RandomValue(remainCount-i, remainAmount)
		remainAmount -= curAmount
		sumAmount += curAmount
	}

	Convery("compare target value if equal source value", t, func() {
		So(sumAmount, ShouldEqual, amount)
	})
}
