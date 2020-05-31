package main

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/wangrenyi/kmpgo/connector"
	"github.com/wangrenyi/kmpgo/kubernetes"
	"github.com/wangrenyi/kmpgo/log"
	"github.com/wangrenyi/kmpgo/proto/pod"
	"golang.org/x/net/netutil"
	"google.golang.org/grpc"
)

func init() {

}

type KmpServer struct {
	grpcServer *grpc.Server
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.DPanicf("main", "server crash: %v", err)
			log.DPanicf("main", "stack: %v", string(debug.Stack()))
		}
	}()

	flag.Parse()

	server := NewServer()
	if err := server.Server(); err != nil {
		log.Errorf("main", "server fail to start: %v", err)
	}
}

func NewServer() *KmpServer {
	grpcServer := grpc.NewServer(grpc.MaxConcurrentStreams(64))

	kubeClient, err := kubernetes.GetKubernetesClient()
	if err != nil {
		panic(err)
	}
	pod.RegisterPodInterfaceServer(grpcServer, &connector.KmpNodeServer{Client: kubeClient})

	return &KmpServer{grpcServer: grpcServer}
}

func (kmpServer *KmpServer) Server() error {
	//TODO config
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Errorf("kmp server failed to listen: %v", err)
		return err
	}
	listen = netutil.LimitListener(listen, 1024)

	//监听退出信号
	go waitToExit(listen)

	//启动服务
	if err := kmpServer.grpcServer.Serve(listen); err != nil {
		log.Errorf("kmp server error: %v", err)
		return err
	}
	return nil
}

func waitToExit(listen net.Listener) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan)

	for {
		s := <-signalChan
		log.Errorf("waitToExit", "get signal: %s", s)
		if s == syscall.SIGINT || s == syscall.SIGQUIT || s == syscall.SIGTERM {
			listen.Close()
			os.Exit(1)
		}
	}
}
