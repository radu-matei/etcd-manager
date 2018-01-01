package main

import (
	"fmt"
	"log"
	"os"

	"k8s.io/client-go/tools/clientcmd"

	"github.com/coreos/etcd-operator/pkg/client"
	"github.com/radu-matei/etcd-manager/pkg/cluster"
)

func main() {
	kubeconfig := getEnvVarOrExit("KUBECONFIG")
	//opImage := "gcr.io/coreos-k8s-scale-testing/etcd-operator"
	//ns := "default"

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("cannot get Kubernetes config: %v", err)
	}

	etcdClient := client.MustNew(config)
	c := cluster.NewCluster("etcd-kube-toolkit", 3)

	res, err := cluster.DeployCluster(etcdClient, c, "default")
	if err != nil {
		log.Fatalf("cannot deploy cluster: %v", err)
	}

	fmt.Printf("created cluster: %v", res.Name)
}

func getEnvVarOrExit(varName string) string {
	value := os.Getenv(varName)
	if value == "" {
		log.Fatalf("missing environment variable %s\n", varName)
	}

	return value
}
