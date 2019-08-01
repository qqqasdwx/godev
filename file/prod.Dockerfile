FROM golang:1.12.4-alpine3.9 AS builder
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache \ 
    git \
    tzdata \
    gcc \
    g++ && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo Asia/Shanghai > /etc/timezone && \
    apk del tzdata

WORKDIR /{{.ProjectName}}

COPY . .

RUN GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /bin/{{.ProjectName}} -gcflags "-N -l" main.go

FROM registry.cn-beijing.aliyuncs.com/intwallyun/base

COPY --from=builder /bin/{{.ProjectName}} /bin/{{.ProjectName}}/{{.ProjectName}}

# ADD conf/app.conf /bin/{{.ProjectName}}/conf/app.conf

CMD [ "/bin/{{.ProjectName}}/{{.ProjectName}}" ]