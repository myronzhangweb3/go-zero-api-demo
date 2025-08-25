# Go Zero API Demo

Go zero API + Postgres SQL

## Init

初始化项目
```bash
goctl api new go-zero-api-demo
```

初始化数据库(PostgresSQL)
```bash
goctl model pg datasource --url "postgres://postgres:123456@127.0.0.1:5432/go-zero-api-demo?sslmode=disable" --table account --dir "./internal/model"
```

初始化API
```bash
goctl api go -api demo.api -dir .
```

## Run

### Go

```bash
go run demo.go
```

### Docker

构建 Dockerfile
```bash
goctl docker --go demo.go --exe go-zero-api-demo
```

构建镜像
```bash
docker build -t go-zero-api-demo:v1 .
```

启动服务
```bash
docker run --rm -it -p 8888:8888 go-zero-api-demo:v1
```


## 测试连接

```bash
curl -i http://localhost:8888/from/you
```