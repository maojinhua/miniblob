## 测试 TLS 通信功能

## 1. 生成 CA 证书文件

在以 TLS 通信模式启动服务器时，需要先生成 CA 证书文件，通过执行以下命令来生成并安装 CA 证书文件：

```bash
# 切换到对应的分支，或者 master 分支
$ git checkout feature/s26

# 生成 CA 证书
$ make ca

# 复制证书文件到配置目录
# $ cp -a _output/cert/ $HOME/.miniblog/
```

## 2. 配置服务器文件

配置 `$HOME/.miniblog/mb-apiserver.yaml` 文件，启用 Gin 服务器模式，并开启 TLS 通信模式、设置 TLS 服务器监听端口，配置如下：

```yaml
server-mode: gin

# 安全服务器相关配置
tls:
  # 是否启动 HTTPS 安全服务器
  # 生产环境建议开启 HTTPS，并使用 HTTPS 协议访问 API 接口
  use-tls: true
  # 证书
  cert: _output/cert/cert/server.crt
  # 证书 Key 文件
  key: _output/cert/cert/server.key

# HTTP 服务器相关配置
http:
  # HTTP 服务器监听地址
  addr: :5555
```

<!-- **注意**：需要将 `tls.cert` 和 `tls.key` 中的 `$HOME`，替换为当前用户的主目录路径。配置文件识别不了 `$HOME` 这种路径表示方式。 -->

上述配置中，开启后的 HTTPS 服务器监听端口为 5555。

## 3. 编译并启动服务器

配置完配置文件后，便可以执行以下命令编译并启动服务器：

```bash
$ make build
$ _output/mb-apiserver
```

启动成功后，日志输出示例如下：
```
{"level":"info","timestamp":"2025-02-01 17:56:46.761","caller":"server/http_server.go:43","message":"Start to listening the incoming requests","protocol":"https","addr":":5555"}
```

## 4. TLS 通信测试

从输出的日志可以知道服务器以 TLS 通信模式监听在 5555 端口上。

新打开一个 Linux 终端，并执行以下操作：

### 测试 1：通过 HTTP 协议访问 /healthz（协议不匹配）
```bash
$ curl http://127.0.0.1:5555/healthz
```
**返回结果：**
```
Client sent an HTTP request to an HTTPS server.
```

### 测试 2：通过 HTTPS 协议访问 /healthz（不指定根证书）
```bash
$ curl https://127.0.0.1:5555/healthz
```
**返回错误：**
```
curl: (60) SSL certificate problem: EE certificate key too weak
More details here: https://curl.haxx.se/docs/sslcerts.html

curl failed to verify the legitimacy of the server and therefore could not establish a secure connection to it. To learn more about this situation and how to fix it, please visit the web page mentioned above.
```

### 测试 3：使用根证书认证服务端
```bash
$ curl https://127.0.0.1:5555/healthz --ciphers DEFAULT@SECLEVEL=1 --cacert $HOME/.miniblog/cert/ca.crt
```
**返回结果：**
```
{"timestamp":"2025-02-01 17:57:37"}
```

### 测试 4：跳过 SSL 检测
```bash
$ curl https://127.0.0.1:5555/healthz -k
```
**返回结果：**
```
{"timestamp":"2025-02-01 17:57:52"}
```

## 总结

该配置文档完整展示了如何：
1. 生成和配置 TLS/SSL 证书
2. 设置服务器以 HTTPS 模式运行
3. 使用不同方式进行 TLS 连接测试
4. 解决证书验证相关问题

代码位于 `feature/s26` 分支的 `scripts/test_tls.sh` 文件中。