# 指定基础镜像
FROM golang:1.14.6 as builder

WORKDIR /go/src/api
ADD ./api ./
RUN CGO_ENABLED=0 go build -ldflags '-d -w -s'

RUN ls -l

# 构建第二阶段
FROM alpine:latest

ENV APP_DIR /go/src/api/
WORKDIR $APP_DIR

COPY --from=builder /go/src/api/api $APP_DIR
COPY --from=builder /go/src/api/config/ ${APP_DIR}config
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN mkdir -p runtime/logs

CMD ./api

EXPOSE 8803