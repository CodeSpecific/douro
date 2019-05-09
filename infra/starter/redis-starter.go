package starter

import (
	"fmt"
)

func init(){
	Register(new(redisStarter))
}

// redisStarter Redis 启动器
type redisStarter struct{
	BaseStarter
}

func(r *redisStarter)Init(ctx StarterContext){
	fmt.Println("redis 已经启动")
}






