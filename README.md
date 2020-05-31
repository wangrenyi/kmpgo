# kmp-go
* kubernetes manage platform -- go client

# 启动
* go build server/kmp_server.go

# 测试
* go test -v test/pod_test.go

# 配置
* 1. 环境  go version go1.13.11 windows/amd64
* 2. VirtualBox 本地虚拟机 -- k8s master 节点
* 3. $HOME/.kube/config 配置  -- 在 windows 上打开 git 命令行,下载 k8s master 节点证书信息到本地环境
* 步骤 0> cd $HOME/.kube
* 步骤 1> sftp 192.168.0.115@k8s-master
* 步骤 2> get /etc/kubernetes/admin.conf config
* 步骤 3> 修改配置文件 config 的 server 属性 为 k8s master 节点 IP 地址
