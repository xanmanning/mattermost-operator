package apis

import (
	miniov1 "github.com/minio/minio-operator/pkg/apis/operator.min.io/v1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, miniov1.SchemeBuilder.AddToScheme)
}
