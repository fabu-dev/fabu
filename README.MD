
# fabu.dev
移动端应用发布平台，用于iOS及Android内测应用的分发。https://fabu.dev



# 生成API接口文档
```bash
# 生成swag文档
go get -u github.com/swaggo/swag/cmd/swag

cd ./api
swag init -o ./api/docs/swagger

```
访问：127.0.0.1:8803/swagger/index.html 查看(ip根据实际情况修改)


# 端口说明
| 端口 | 说明 |
| :-- | :-- |
| 8803 | api端口 |
| 8804 | web端口 |
| 8805 | mysql端口 |
| 8806 | redis端口 |