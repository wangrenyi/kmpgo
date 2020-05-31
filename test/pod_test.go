package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/wangrenyi/kmpgo/proto/common"

	convey "github.com/smartystreets/goconvey/convey"
	"github.com/wangrenyi/kmpgo/log"
	"github.com/wangrenyi/kmpgo/model"
	pb "github.com/wangrenyi/kmpgo/proto/pod"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	appsv1 "k8s.io/api/apps/v1"
)

const (
	address string = "127.0.0.1"
	port    string = "9090"
)

func TestConveyCase_PodGet(t *testing.T) {
	convey.Convey("PodGet - 查询参数 Pod 详情", t, func() {
		conn, err := grpc.Dial(address+":"+port, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Errorf("did not connect: %v", err)
		}
		defer func() {
			err = conn.Close()
			if err != nil {
				log.Errorf("close connection fail: %v", err)
			}
		}()

		client := pb.NewPodInterfaceClient(conn)

		requestData, err := json.Marshal(&model.DeploymentListReq{})
		if err != nil {
			log.Errorf("reData json.Marshal error: %v", err)
		}

		result, err := client.PodList(context.Background(), &common.InvokeServiceRequest{
			Data: string(requestData),
		})
		if err != nil {
			log.Errorf("client podGet error: %v", err)
		}
		log.Infof("PodGet response: %v", result)
		fmt.Printf("PodList response: %v \n", result)

		response := new(appsv1.Deployment)
		err = json.Unmarshal([]byte(result.Data), response)
		if err != nil {
			log.Errorf("client podGet result Unmarshal error: %v", err)
		}

		convey.Convey("err should be nil ", func() {
			convey.So(err, convey.ShouldBeNil)
		})

	})
}
