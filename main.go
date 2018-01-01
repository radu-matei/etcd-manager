package main

import (
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	api "github.com/coreos/etcd-operator/pkg/apis/etcd/v1beta2"
	"github.com/coreos/etcd-operator/pkg/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	kubeconfig := getEnvVarOrExit("KUBECONFIG")
	//opImage := "gcr.io/coreos-k8s-scale-testing/etcd-operator"
	//ns := "default"

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("cannot get Kubernetes config: %v", err)
	}
	cli, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("cannot get Kubernetes client: %v", err)
	}

	etcdClient := client.MustNew(config)

}

func newCluster(name string, size int) *api.EtcdCluster {
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

func getEnvVarOrExit(varName string) string {
	value := os.Getenv(varName)
	if value == "" {
		log.Fatalf("missing environment variable %s\n", varName)
	}

	return value
}
