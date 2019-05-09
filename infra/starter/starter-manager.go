package starter

// starterManager 启动器管理器
type starterManager struct{
	Starters []Starter
}

// Register 注册启动器
func(m *starterManager)Register(starter Starter){
	m.Starters=append(m.Starters,starter)
}

// GetAllStarters 获取所有的启动器
func (m *starterManager) GetAllStarters()[]Starter{
	return m.Starters
}

// Run 运行所有的启动器
func(m *starterManager)Run(){
	var starterContext = make(StarterContext)
	//初始化所有的启动器
	for _,starter:=range m.Starters{
		starter.Init(starterContext)
	}
	//建立所有启动器的环境
	for _,starter:=range m.Starters{
		starter.SetUp(starterContext)
	}
	//启动所有的启动器
	for _,starter:=range m.Starters{
		starter.Start(starterContext)
	}
}

// StarterManager 全局的启动器
var StarterManager = new(starterManager)

// Register 全局注册启动器的方法
func Register(starter Starter){
	StarterManager.Register(starter)
}









