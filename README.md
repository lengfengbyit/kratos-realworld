# Kratos Project Template

## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

## realworld
[realworld-api-docs](https://realworld-docs.netlify.app/specifications/backend/endpoints/)

## 开发流程
1. 在 api 目录下编写 *.proto 文件
2. 执行 make api 生成 api 文件
3. 执行 kratos proto server api/server/*.proto -t internal/service 命令创建 service 文件
4. 在 service 文件中开发业务逻辑
5. 将初始化函数添加的 service/service.go 的 wire ProviderSet 中
6. 添加对应的 biz 和 data 文件，并将初始化函数需要添加到 wire ProviderSet 
7. 添加 ent schema 文件,生成数据库操作代码：ent new <table_name>, table_name 用大驼峰命名
8. 在生成的 data/end/schema/<table_name>.go 文件中添加字段、索引等
9. 完成 service, biz, data 等文件下的代码编写
10. 将代码注册到 serve 中： 在 server/http.go 文件注册当前服务
11. 运行 make wire 完成依赖注入