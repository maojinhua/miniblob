Biz层依赖Store层，所以实现了Store层代码之后，便可以实现Biz层代码。Biz层代码，主要用来实现系统中REST资源的各类业务操作，例如用户资源的增删改查等。

整个miniblog项目的设计较为规范，规范化的项目设计带来的优点之一是开发方式的一致性和开发效率的提升。miniblog项目Biz层的开发方式与Store层的开发方式保持一致。

API接口定义
在开发Biz层代码之前，需要先定义好API接口的请求入参和返回参数。为此，新建了pkg/api/apiserver/v1/user.proto文件和pkg/api/apiserver/v1/post.proto文件，分别保存了用户接口和博客接口的请求入参和返回参数。具体接口定义见feature/s18分支pkg/api/apiserver/v1/目录中的对应文件。

因为请求入参和返回参数（例如CreateUserRequest和CreateUserResponse）会提供给接口调用方（客户端），所以需要将接口定义保存在pkg/api目录下。另外，考虑到未来miniblog可能会加入多个服务，每个服务都有自己的API定义，miniblog项目选择了将每个服务的API定义保存在独立的服务目录下，例如pkg/api/apiserver。

考虑到未来API接口的版本升级，miniblog项目将接口进行了版本化处理，v1版本的接口保存在pkg/api/apiserver/v1目录下，v2版本的接口保存在pkg/api/apiserver/v2目录下。

在Go项目开发中，开发者在实现新功能开发时，要时刻考虑到功能未来的扩展能力。否则，未来功能需要升级或扩展时，将面临大量的代码重构，而且重构成本甚至比新开发成本还要高。

在定义完接口之后，还需要执行以下命令，编译Protobuf文件：

```
 make protoc
```
IBiz接口定义及实现
Biz层代码保存internal/apiserver/biz/biz.go文件中，接口名为IBiz，定义如下：