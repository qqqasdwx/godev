FROM registry.cn-beijing.aliyuncs.com/intwallyun/base:latest
COPY ./myapp /bin/{{.ProjectName}}/{{.ProjectName}}
WORKDIR /bin/{{.ProjectName}}
CMD [ "./{{.ProjectName}}" ]
