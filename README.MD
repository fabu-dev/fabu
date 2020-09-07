
# fabu.dev
移动端应用发布平台，用于iOS及Android内测应用的分发。https://fabu.dev



# 生成API接口文档
```bash
# 生成swag文档
go get -u github.com/swaggo/swag/cmd/swag
swag init -o ./api/docs/swagger
```