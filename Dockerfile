# 构建
FROM golang:1.23-alpine as builder
WORKDIR /home/project
ENV GOPROXY=https://goproxy.cn
COPY ./ ./
RUN go mod tidy && \
    go mod download && \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add tzdata
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o k8s-deploy

# 打包
FROM alpine as runner
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /home/project/k8s-deploy /home/project/
COPY --from=builder /home/project/etc /home/project/etc
COPY --from=builder /home/project/asset/swagger/ /home/project/asset/swagger
COPY --from=builder /home/project/ui/ /home/project/ui
WORKDIR /home/project
ENTRYPOINT ["./k8s-deploy"]
CMD ["-env", "prod"]
