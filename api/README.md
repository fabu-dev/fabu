
# go-web-frame
使用gin构建的web项目开发框架，采用web开发中典型分层架构-MVC, 对于view层，建议采用前后端分离的方式实现，另外增加过滤层和业务逻辑处理层。


# 结构
├── config   
├── controller   
├── docs    
├───── swagger  
├── filter  
├── model  
├── service    
├── pkg   
├── router   
├── go.mod  
├── main.go   
├── README.md   


# Feature
+ RESTful API
+ Gorm
+ Go-redis
+ Yaml config file
+ Gin
+ pprof
+ swagger



# 其他
```bash
# 生成swag文档
go get -u github.com/swaggo/swag/cmd/swag
swag init -o ./docs/swagger
```