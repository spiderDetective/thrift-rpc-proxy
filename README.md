使用方法：
1. 注册rpc服务
在consul中注册：urlprefix-/{thrift rpc method name}  proto=thrift 
使用proto=thrift 指明是thrift协议，和grpc区分开。

2. 查看
服务向fabio注册成功后，可以通过fabio的面板查看。

3. 部署
启动fabio时需要指定thrift rpc端口地址。例如：
-proxy.addr=:80,:443;cs=some-name,:12345;proto=thrift