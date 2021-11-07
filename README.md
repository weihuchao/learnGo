## 1 基础入门和编程思想

### 1.1 准备工作

- 官网：https://golang.org/

- 国内下载：https://studygolang.com/dl

- 国内镜像：https://goproxy.cn/
    - 设置代理
    - ```bash
$ go env -w GO111MODULE=on 
$ go env -w GOPROXY=https://goproxy.cn,direct
    ```
- IDE

    - Goland/Idea + Go插件
        - 收费
        - 语法支持非常好
        - 安装vgo：http://plugins.jetbrains.com/plugin/10753-vgo-support/versions
        - 最新版默认集成了
    - VScode
        - 免费
        - 全栈开发的时候更高效

- 新建一个go项目

    - 选择Go Modules这一栏；
    - `go.mod`文件会声明一下包名，一般都会带上域名；
    - 新建go文件可以选择一个`simple application`里面会自动新增一个`main`函数；
    - 可以安装一个插件叫做：`file wathchers`，可以自动添加引用和自动格式化文件；
    - 可以设置快捷键：
        - 删除一行（delete line）；
        - 向前（forward）；
        - 后退（back）；
        - 跳转到定义（go to declaration）；
    - 运行文件旁边有个设置配置，可以设置`run kind`来保证运行的是文件而不是package；

### 1.2 基础部分

- 变量
    - 定义、使用
    - 包变量、局部变量
    - 常量
    - 枚举
- 变量类型
    - 基本类型
    - 最大最小值
    - string、byte、rune
    - float
        - 精度问题
    - 类型转换
    - string的操作
- 条件循环
    - if
    - switch
    - for
- 函数
- 指针
    - 值传递
    - 引用传递
- 数组
    - 定义
    - 修改
    - 值传递
    - 遍历
- 切片
    - 定义、本质、使用
    - 多次切片
    - 切片的cap
    - append
    - copy
- map
    - 创建
    - 获取值
    - 修改、删除
    - 遍历
- 结构体
- 依赖管理
    - gopath
    - vendor
    - go mod
        - 环境变量的设置
        - 拉取依赖
        - 更改依赖的版本
- 项目结构
    - 包
    - main函数、main包
        - 同一个目录下不能有多个main包
    - 名称定义，接口体的名字可以省略包名，如`tree.TreeNode`可以改为`tree.Node`
    - 检查项目编译是否通过`go build ./...`
    - 编译生成文件`go install ./...`
        - 生成的文件在`${GOPATH}/bin`
        - 生成的文件可以直接执行
- 附加部分
    - 最大子串长度
    - 欧拉公式
    - 读文件
    - 数字转为二进制
    - 快捷键：
        - 快速返回值：win + shift + l
        - 快速重命名：选定之后 win + shift + l
        - 实现接口：win + shift + p
        - 搜索文档接口：ctrl + o
        -  