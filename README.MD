
# fabu.dev
移动端应用发布平台，用于iOS及Android内测应用的分发。https://fabu.dev



# 生成API接口文档
```bash
# 生成swag文档
go get -u github.com/swaggo/swag/cmd/swag

cd ./api
swag init -o ./docs/swagger

```

# 后端项目启动
```bash
cd ./api

go run main.go
```
访问：127.0.0.1:8803/swagger/index.html 查看(ip根据实际情况修改)

# 前端项目启动
```bash
yarn install

yarn run serve
```
可以注册新账号使用，或用演示账号登录（账号和密码均为admin）

# 通过docker部署
```bash
# cd 到根目录

# 构建镜像，并创建容器
docker-compose up -d

# 关闭服务器
docker-compose stop

# 启动容器
docker-compose start

```

# 端口说明
| 端口 | 说明 |
| :-- | :-- |
| 8803 | api端口 |
| 8804 | web端口 |
| 8805 | mysql端口 |
| 8806 | redis端口 |