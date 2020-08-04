
## 格式规范

#### 文件夹名称 和 package包名

######只包含小写字母a-z，如 api, cmd, githook

#### 代码文件名称

######只包含小写字母a-z和下划线_，如 [rpc_client.go](https://gitee.com/Skyd188/micro_services/blob/master/internal/app/client/myendpoint/rpc_client.go)

#### 函数/方法名称

######大小写字母a-z A-Z，通常避免使用特殊符号和数字

```go
func MakeEchoEndpoint(svc pb.MyServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.EchoRequest)
		return svc.Echo(ctx, req)
	}
}
```

#### 常量定义

######驼峰命名，且通常在文件开头统一声明

```go
package user

import "github.com/go-kit/kit"

const (
    userName    = "Tom"
    companyName = "Meross"
)
```

#### 变量定义

######驼峰命名，且函数内变量应尽量在函数开头统一声明

```go
func DoSomeThing(str string) {
    var (
        a, b, c string
        x, y, z int
    )

    //something
    return
}
```

## 目录说明

### `/api`

模板文件，第三方库所需的数据文件，JSON，proto定义等

### `/cmd`

各个项目模块的[`main`]函数入口程序

### `/config`

项目配置文件，监听端口，数据库信息等

### `/githook`

githook脚本文件

### `/init`

读取配置文件到全局变量，初始化项目信息

### `/internal`

项目内部引用的包

    /app 对应cmd目录，只被对应模块引用
    /pkg 模块公用代码
    
### `/tools`

独立于业务之外的工具函数


