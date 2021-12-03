package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"k8s.io/client-go/rest"
	k8sclientapiv1 "k8s.io/client-go/tools/clientcmd/api/v1"
	"sigs.k8s.io/yaml"
)

func main() {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	caData, err := ioutil.ReadFile(cfg.TLSClientConfig.CAFile)
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	url := fmt.Sprintf("http://127.0.0.1:%s", port)
	c, err := GetKubeconfigWithSAToken("local", "local", url, caData, cfg.BearerToken)
	if err != nil {
		panic(err)
	}

	filename := os.Getenv("KUBECONFIG_FILENAME")
	if filename == "" {
		filename = "/kubeconfigs/kubeconfig"
	}

	err = ioutil.WriteFile(filename, []byte(c), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("kubeconfig successfully written to '%s'\n", filename)
}

func GetKubeconfigWithSAToken(name, username, url string, caData []byte, saToken string) (string, error) {
	config := k8sclientapiv1.Config{
		APIVersion: k8sclientapiv1.SchemeGroupVersion.Version,
		Kind:       "Config",
		Clusters: []k8sclientapiv1.NamedCluster{
			{
				Name: name,
				Cluster: k8sclientapiv1.Cluster{
					CertificateAuthorityData: caData,
					Server:                   url,
				},
			},
		},
		Contexts: []k8sclientapiv1.NamedContext{
			{
				Name: name,
				Context: k8sclientapiv1.Context{
					Cluster:  name,
					AuthInfo: username,
				},
			},
		},
		CurrentContext: name,
		AuthInfos: []k8sclientapiv1.NamedAuthInfo{
			{
				Name: username,
				AuthInfo: k8sclientapiv1.AuthInfo{
					Token: saToken,
				},
			},
		},
	}

	y, err := yaml.Marshal(config)
	if err != nil {
		return "", err
	}

	return string(y), nil
}
