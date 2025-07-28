## miniblog 项目
### 编译命令
make build
### 启动命令
_output/mb-apiserver -c $HOME/.miniblog/mb-apiserver.yaml

### grpc-gateway
#### ​​基本概念​​
grpc-gateway 是 protoc 工具的插件，用于将 gRPC 服务映射到 HTTP/REST API。
​​前置条件​​
#### 需确保系统已安装：
protoc 工具（Protocol Buffers 编译器）
#### ​​必备插件​​
安装以下两个核心插件：
​​protoc-gen-grpc-gateway​​
功能：为 gRPC 服务生成 HTTP/REST API 反向代理代码
作用：实现对 gRPC 服务的 HTTP 映射支持

​​protoc-gen-openapiv2​​
功能：从 Protobuf 描述中生成 OpenAPI v2 (Swagger) 定义文件
作用：生成标准化的 API 文档

​​#### 安装命令​​
```
安装 gRPC Gateway 插件
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.24.0

安装 OpenAPI 文档生成插件
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.24.0
```
