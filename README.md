# douro

代号：斗罗，又名：红包系统(Red Envelop System)，是高并发资金交易系统的一种表现形式。

## 在线演示

在线演示地址: http://106.12.130.92/

## 环境准备

- [ ] Go 1.11.1 

### 开发框架

- [ ] Web开发框架：https://github.com/kataras/iris
- [ ] Api开发框架：https://github.com/gin-gonic/gin
- [ ] Api文档框架：https://github.com/swaggo/gin-swagger
- [ ] 日志框架：https://github.com/sirupsen/logrus
- [ ] 验证框架：https://github.com/go-playground/validator
- [ ] ORM框架：https://github.com/go-xorm/xorm
- [ ] 配置文件框架：https://github.com/rickar/props
- [ ] 自动化测试框架：https://github.com/smartystreets/goconvey

### FAQ
*Q*: 无法访问类似于`golang.org/x/tools`这样的依赖包，怎么办?

系统使用`Go Module`管理依赖，当运行`build/run`的时候，没有科学上网会出现类似于`golang.org/x/tools`无法访问到的情况，建议配置`GoProxy (https://goproxy.io)`
