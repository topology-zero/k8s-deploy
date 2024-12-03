### k8s-deploy 后台管理

在 k8s 中部署 Deployment、Service、IstioGateway、IstioVirtualService、IstioDestinationRule

### 目前支持的 CDR

* Deployment
* ReplicationController
* ReplicaSet
* StatefulSet
* DaemonSet
* Pod
* Job
* CronJob
* Service
* Endpoint
* PersistentVolumeClaim
* ConfigMap
* IstioGateway
* IstioVirtualService
* IstioDestinationRule

### 特点

无需登录服务器使用命令 `kubectl apply -f xxx.yaml`，只需动动鼠标点点就可以实现 k8s 部署 Deployment、Service 等

### 部署方法

1. 创建 MYSQL 数据库，并将项目下的 `k8s-deploy.sql` 导入
2. 修改 `k8s-deploy.yaml` 的环境变量配置，配置文件在项目下的 `etc/xxx.yaml` 目录，可以使用环境变量覆盖
3. 将项目下 k8s-deploy.yaml 复制到 k8s 服务器并部署 `kubectl apply -f k8s-deploy.yaml`

### 注意事项

默认账号密码 

```yaml
admin
123456
```

> [!CAUTION]
> 部署成功第一件事：立即修改账号密码.

### 联系

##### 可直接在 issues 提出

##### 微信

<img decoding="async" src="http://tc.masterjoy.top/%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20230216101038.jpg" width="30%" />
