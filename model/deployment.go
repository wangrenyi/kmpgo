package model

type DeploymentCreateReq struct {
	Namespace string            `json:"namespace"`
	Name      string            `json:"name"`
	Replicas  int32             `json:"replicas"`
	Labels    map[string]string `json:"labels"`
}

type DeploymentGetReq struct {
	Name string `json:"name"`
}

type DeploymentListReq struct {
}
