# cncamp-module12
CNCAMP module 12 homework - Istio

## 重构
- 基于[Go 面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)重构了代码结构
- 使用gin web framework替换了httprouter
- 使用[gin Prometheus中间件](https://github.com/penglongli/gin-metrics)

## 作业
- 把我们的 httpserver 服务以 Istio Ingress Gateway 的形式发布出来
- 如何实现安全保证
- 七层路由规则
- 考虑 open tracing 的接入

### 实现
- 删除了老的nginx ingress
- 新增[istio 配置](deployments/manifests/istio.yaml)
- 更改[certificate配置](deployments/manifests/certificate.yaml)，重新生成证书在istio-system命名空间，供默认istio ingress gateway使用
- 安装了Jaeger作为tracing监控

![Alt text](assests/img/cncamp.jpg?raw=true "hello")

![Alt text](assests/img/grafana.jpg?raw=true "grafana")

![Alt text](assests/img/jaeger.jpg?raw=true "jaeger")