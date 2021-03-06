package test

import (
	"testing"

	kubeClient "github.com/containerum/kube-client/pkg/client"
	"github.com/containerum/kube-client/pkg/rest/re"
)

func newClient(test *testing.T) *kubeClient.Client {
	client, err := kubeClient.NewClient(
		kubeClient.Config{
			RestAPI: re.NewResty(re.SkipTLSVerify, re.WithHost("https://192.168.88.200:8082")),
			User: kubeClient.User{
				Role: "user",
			},
		})
	if err != nil {
		test.Fatalf("error while creating client: %v", err)
	}
	return client
}
