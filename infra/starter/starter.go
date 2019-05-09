package starter

// StarterContext 启动器的上下文
type StarterContext map[string]interface{}

// Starter 启动器接口
type Starter interface {
	Init(StarterContext) 
	SetUp(StarterContext)
	Start(StarterContext)
	StartBlocking() bool
	Stop(StarterContext)
}

// 判断是BaseStarter是否实现了接口Starter，以帮助在程序编译的时候检查
var _ Starter=new(BaseStarter)

// BaseStarter 基础启动器
type BaseStarter struct{}

// Init 初始化
func(s *BaseStarter)Init(ctx StarterContext){}
// SetUp 建立环境
func(s *BaseStarter)SetUp(ctx StarterContext){}
// Start 启动
func(s *BaseStarter)Start(ctx StarterContext){}
// StartBlocking 启动是否阻塞
func(s *BaseStarter)StartBlocking()bool{return false}
// Stop 优雅的停止
func(s *BaseStarter)Stop(ctx StarterContext){}
