package connector

import (
	"github.com/wangrenyi/kmpgo/kubernetes"
	"github.com/wangrenyi/kmpgo/log"
	"github.com/wangrenyi/kmpgo/model"
	common "github.com/wangrenyi/kmpgo/proto/common"
	"github.com/wangrenyi/kmpgo/util"
	"golang.org/x/net/context"
)

type KmpPodServer struct {
	Client *kubernetes.KubernetesClient
}

func (kmpServer *KmpPodServer) PodCreate(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {
	podCreateReq := &model.DeploymentCreateReq{}
	if err := util.Unmarshal(req.Data, podCreateReq); err != nil {
		log.Errorf("PodCreate", "req data Unmarshal", "error: %v", err)
		return nil, err
	}
	result, err := kmpServer.Client.DeploymentCreate(ctx, podCreateReq)
	if err != nil {
		log.Errorf("PodCreate", "DeploymentCreate", "error: %v", err)
		return nil, err
	}
	return &common.InvokeServiceResponse{Data: model.GenerateJSON(model.Success, result)}, nil
}

func (kmpServer *KmpPodServer) PodUpdate(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpPodServer) PodUpdateStatus(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpPodServer) PodDelete(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpPodServer) PodDeleteCollection(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpPodServer) PodGet(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {
	podGetReq := &model.DeploymentGetReq{}
	if err := util.Unmarshal(req.Data, podGetReq); err != nil {
		log.Errorf("PodGet", "req data Unmarshal", "error: %v", err)
		return nil, err
	}
	result, err := kmpServer.Client.DeploymentGet(ctx, podGetReq)
	if err != nil {
		log.Errorf("PodGet", "DeploymentGet", "error: %v", err)
		return nil, err
	}
	return &common.InvokeServiceResponse{Data: model.GenerateJSON(model.Success, result)}, nil
}

func (kmpServer *KmpPodServer) PodList(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {
	podListReq := &model.DeploymentListReq{}
	if err := util.Unmarshal(req.Data, podListReq); err != nil {
		log.Errorf("PodList", "req data Unmarshal", "error: %v", err)
		return nil, err
	}
	result, err := kmpServer.Client.DeploymentList(ctx, podListReq)
	if err != nil {
		log.Errorf("PodList", "DeploymentList", "error: %v", err)
		return nil, err
	}
	return &common.InvokeServiceResponse{Data: model.GenerateJSON(model.Success, result)}, nil
}

func (kmpServer *KmpPodServer) PodWatch(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpPodServer) PodPatch(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpPodServer) PodGetEphemeralContainers(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpPodServer) PodUpdateEphemeralContainers(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

// pod expansion
func (kmpServer *KmpPodServer) PodBind(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpPodServer) PodEvict(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpPodServer) PodGetLogs(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}
