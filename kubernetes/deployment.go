package kubernetes

import (
	"github.com/wangrenyi/kmpgo/model"
	"golang.org/x/net/context"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (client *KubernetesClient) DeploymentCreate(ctx context.Context, req *model.DeploymentCreateReq) (*appsv1.Deployment, error) {
	deploymentsClient := client.clientSet.AppsV1().Deployments(corev1.NamespaceDefault)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &req.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: req.Labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: req.Labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	return deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
}

func (client *KubernetesClient) DeploymentGet(ctx context.Context, req *model.DeploymentGetReq) (*appsv1.Deployment, error) {
	deploymentsClient := client.clientSet.AppsV1().Deployments(corev1.NamespaceDefault)
	return deploymentsClient.Get(ctx, req.Name, metav1.GetOptions{})
}

func (client *KubernetesClient) DeploymentList(ctx context.Context, req *model.DeploymentListReq) (*appsv1.DeploymentList, error) {
	deploymentsClient := client.clientSet.AppsV1().Deployments(corev1.NamespaceDefault)
	return deploymentsClient.List(ctx, metav1.ListOptions{})
}
