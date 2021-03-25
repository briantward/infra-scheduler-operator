module github.com/briantward/infra-scheduler-operator

go 1.15

require (
	github.com/go-logr/logr v0.3.0
	github.com/kevinburke/go-bindata v3.22.0+incompatible
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/openshift/machine-config-operator v0.0.1-0.20210324214351-558c30a44597
	k8s.io/api v0.20.0
	k8s.io/apimachinery v0.20.0
	k8s.io/client-go v0.20.0
	sigs.k8s.io/controller-runtime v0.7.2
)
