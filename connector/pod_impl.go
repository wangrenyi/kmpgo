package connector

import (
	"github.com/wangrenyi/kmpgo/kubernetes"
	"github.com/wangrenyi/kmpgo/log"
	"github.com/wangrenyi/kmpgo/model"
	common "github.com/wangrenyi/kmpgo/proto/common"
	"github.com/wangrenyi/kmpgo/util"
	"golang.org/x/net/context"
)

type KmpNodeServer struct {
	Client *kubernetes.KubernetesClient
}

func (kmpServer *KmpNodeServer) PodCreate(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {
	podCreateReq := &model.DeploymentCreateReq{}
	if err := util.Unmarshal(req.Data, podCreateReq); err != nil {

	}
	result, err := kmpServer.Client.DeploymentCreate(ctx, podCreateReq)
	if err != nil {
		log.Errorf("PodCreate", "DeploymentCreate", "error: %v", err)
		return nil, err
	}
	return &common.InvokeServiceResponse{Data: model.GenerateJSON(model.Success, result)}, nil
}

func (kmpServer *KmpNodeServer) PodUpdate(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpNodeServer) PodUpdateStatus(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpNodeServer) PodDelete(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpNodeServer) PodDeleteCollection(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpNodeServer) PodGet(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {
	podGetReq := &model.DeploymentGetReq{}
	if err := util.Unmarshal(req.Data, podGetReq); err != nil {

	}
	result, err := kmpServer.Client.DeploymentGet(ctx, podGetReq)
	if err != nil {
		log.Errorf("PodGet", "DeploymentGet", "error: %v", err)
		return nil, err
	}
	return &common.InvokeServiceResponse{Data: model.GenerateJSON(model.Success, result)}, nil
}

func (kmpServer *KmpNodeServer) PodList(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {
	podListReq := &model.DeploymentListReq{}
	if err := util.Unmarshal(req.Data, podListReq); err != nil {

	}
	result, err := kmpServer.Client.DeploymentList(ctx, podListReq)
	if err != nil {
		log.Errorf("PodList", "DeploymentList", "error: %v", err)
		return nil, err
	}
	return &common.InvokeServiceResponse{Data: model.GenerateJSON(model.Success, result)}, nil
}

func (kmpServer *KmpNodeServer) PodWatch(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpNodeServer) PodPatch(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpNodeServer) PodGetEphemeralContainers(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpNodeServer) PodUpdateEphemeralContainers(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

// pod expansion
func (kmpServer *KmpNodeServer) PodBind(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpNodeServer) PodEvict(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}

func (kmpServer *KmpNodeServer) PodGetLogs(ctx context.Context, req *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error) {

	return nil, nil
}
