# `protoc-go-inject-tag` 工具使用说明

## 作用概述

`protoc-go-inject-tag` 是一款用于为 Protobuf 消息体字段添加 Go 结构体标签的工具。其主要解决的是：标准的 `protoc` 编译器生成的 Go 结构体字段默认不包含特定的结构体标签（如 `uri`、`form` 等），而此类标签在 Go Web 开发（例如使用 Gin 等框架进行数据绑定）时常常是必需的。

## 使用方法

### 1. 在 Protobuf 文件中添加注释

在您的 `.proto` 文件中的指定字段上，使用 `// @gotags` 注释来声明需要注入的标签。

```protobuf
// DeleteUserRequest 表示删除用户请求
message DeleteUserRequest {
    // userID 表示用户 ID
    // @gotags uri:"userID"
    string userID = 1;
}
```

### 2. 安装工具

通过以下命令安装 `protoc-go-inject-tag` 工具：

```bash
go install github.com/favadi/protoc-go-inject-tag@latest
```

### 3. 集成到构建流程（Makefile）

通常需要在通过 `protoc` 生成 `.pb.go` 文件后，执行此工具来注入标签。一个高效的做法是将其集成到您的 `Makefile` 中。

```makefile
protoc: # 编译 protobuf 文件。
    ... （您原有的 protoc 生成命令）
    @find . -name "*.pb.go" -exec protoc-go-inject-tag -input={} \;
```

### 4. 执行命令

更新 `Makefile` 后，运行以下命令来重新生成并处理 Go 文件：

```bash
make protoc
```

此流程确保了 Protobuf 文件被编译后，所有的 `.pb.go` 文件会自动被扫描并注入您在注释中定义的 Go 结构体标签。