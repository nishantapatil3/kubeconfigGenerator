module github.com/cisco-app-networking/apioverride

go 1.16

require (
	k8s.io/client-go v1.5.2
	sigs.k8s.io/yaml v1.2.0
)

replace k8s.io/client-go => k8s.io/client-go v0.22.1
