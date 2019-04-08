package algo

import (
	"testing"
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
	if sumAmount == amount {
		t.Log("测试通过了")
	} else {
		t.Errorf("测试失败,输出结果为:%d,预期结果为:%d", sumAmount, amount)
	}
}
