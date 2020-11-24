# gf-api
GF(Go Frame) api 脚手架

启动 + 自动生成 swagger 

```shell
gf run main.go --swagger
```
# gf cli Install

- [安装](https://github.com/gogf/gf-cli)

# swagger
生成 swagger 文件 

```shell
gf pack swagger packed/swagger.go -n packed -y
```
# 目录
```html
    app                         项目路径
        controller              控制器路径      参数接收、判断
        dao                     底层数据提供
        entity                  结构体路径
        logic                   逻辑路径        数据组装
        middleware              中间件路径
        model              
        service                 数据提供
    boot
    common                      
    config
    document
    i18n
    packed
    public
    router
    swagger
```