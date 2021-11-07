## 1 GOPATH

## 2 VENDOR

## 3 GO MOD

这才是真正解决依赖的最好办法.

创建好项目之后, 需要设置以下:

### 3.1 设置环境变量

```bash
go env -w GO111MODULE=on 
go env -w GOPROXY=https://goproxy.cn,direct
```

或者在Goland中, Preference -> GO -> GO MODULES 新增这两个变量

### 3.2 初始化

```bash
➜ go mod init
go: cannot determine module path for source directory /Users/weihuchao/work/study/golang/learngo (outside GOPATH, module path must be specified)

Example usage:
        'go mod init example.com/m' to initialize a v0 or v1 module
        'go mod init example.com/m/v2' to initialize a v2 module

Run 'go help mod init' for more information.
➜ go mod init imooc.com/weihc/learngo
```

也就是在`go mod init`的时候需要指定`package`的名字.

### 3.3 安装依赖包

1. 默认安装

```bash
➜ go get -u go.uber.org/zap
go get: added go.uber.org/atomic v1.9.0
go get: added go.uber.org/multierr v1.7.0
go get: added go.uber.org/zap v1.19.1
```

此时在`go.mod`中会新增:

```
require (
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
)
```

同时`go.mod`下面会出现文件`go.sum`这个记录了具体安装的版本和MD5.

这个`// indirect`表示现在安装了, 还没有被使用

2. 使用之后`// indirect`就会消失, 如果没有可以执行一下`go mod tidy`.
3. 安装的时候可以指定版本, 上github上的`Tags`复制到后面, 前面加上`@`即可.

```bash
➜ go get -u go.uber.org/zap@v1.18.1
```

4. 更改版本之后别的版本不需要的可以用命令清除一下.

```bash
➜ go mod tidy  
```

### 3.4 源文件位置

可以看到, 安装了`zap`之后查找文件就到以下这个目录下:

```
/Users/weihuchao/go/pkg/mod/go.uber.org/zap@v1.18.1/
```

此时找源文件的顺序为:

```
1. /usr/local/go/src/  (也就是: ${GOROOT}/src)
2. /Users/weihuchao/go/pkg/mod/ (也就是: ${GOPATH}/pkg/mod)
```

### 3.5 对已有项目重建go mod

```bash
1. 写入环境变量
2. go mod init xxx
3. go build ./...
```

其中`go build ./...`就会触发当前目录下所有文件的build, 进而自动拉取依赖包(自动`go get`).
