# visible-mq
可视化操作的message queue

## <font color="red">!!!请勿使用在生产环境</font>

本项目旨在通过可视化的操作来进行话题的订阅，在php层面实现无感知的操作
#### 依赖
- Redis
- Go
- PHP
    - PhpRedis扩展

 #### 使用
 1. cd go
 2. go mod download
 3. cd web && go run main.go

 #### 流程图
 ![流程图](https://raw.githubusercontent.com/yankewei/blog/master/image/visible-mq%E6%B5%81%E7%A8%8B%E5%9B%BE.jpg)