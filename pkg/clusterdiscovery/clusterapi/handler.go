package clusterapi

import (
	"github.com/zach593/karmada/pkg/util"
	"github.com/zach593/karmada/pkg/util/informermanager/keys"
)

// ClusterWideKeyFunc generates a ClusterWideKey for object.
func ClusterWideKeyFunc(obj interface{}) (util.QueueKey, error) {
	return keys.ClusterWideKeyFunc(obj)
}
