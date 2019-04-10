package algo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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

	Convey("compare a and b is equal", t, func() {
		So(sumAmount, ShouldEqual, amount)
	})
}
