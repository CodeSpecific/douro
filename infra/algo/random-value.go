package algo

import (
	"math/rand"
	"time"
)

const min int = 100

func init() {
	//设置随机数的种子，防止每次程序重新启动，随机数相同
	rand.Seed(time.Now().UnixNano())
}

/*
 * 随机数算法，只是简单实现业务逻辑，不考虑是否分配均匀
 */
func RandomValue(remainCount, remainAmount int) int {
	if remainCount == 1 {
		return remainAmount
	}
	remainAmount = remainAmount - remainCount*min
	x := rand.Intn(remainAmount) + min
	return x
}
