package cluster

import (
	api "github.com/coreos/etcd-operator/pkg/apis/etcd/v1beta2"
	"github.com/coreos/etcd-operator/pkg/generated/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeployCluster creates a new etcd cluster
func DeployCluster(client versioned.Interface, cluster *api.EtcdCluster, namespace string) (*api.EtcdCluster, error) {
	cluster.Namespace = namespace

	return client.EtcdV1beta2().EtcdClusters(namespace).Create(cluster)
}

// NewCluster returns a new etcd cluster object given a name and size
func NewCluster(name string, size int) *api.EtcdCluster {
	return &api.EtcdCluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       api.EtcdClusterResourceKind,
			APIVersion: api.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: name,
		},
		Spec: api.ClusterSpec{
			Size: size,
		},
	}
}
