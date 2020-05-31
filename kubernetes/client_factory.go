package kubernetes

//import "sync"
//
//var (
//	kubernetesClient *KubernetesClient
//	syncOnce         sync.Once
//)
//
//type ClientFactory interface {
//	GetKubernetesClient() (*KubernetesClient, error)
//}
//
//type clientFactory struct {
//}
//
//func (cf *clientFactory) GetKubernetesClient() (*KubernetesClient, error) {
//	syncOnce.Do(func() {
//		kubernetesClient = NewKubernetesClient()
//	})
//
//	return kubernetesClient, nil
//}
