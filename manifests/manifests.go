package manifests

import (
	"bytes"
	"io"

	corev1 "k8s.io/api/core/v1"

	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

const (
	InfraNamespaceAsset         = "assets/namespace.yaml"
	InfraMachineConfigPoolAsset = "assets/infra-machineconfigpool.yaml"

	OwningInfraLabel = "infrascheduler.bward.dev/owning-infra"
)

func MustAssetReader(asset string) io.Reader {
	return bytes.NewReader(MustAsset(asset))
}

func InfraNamespace() *corev1.Namespace {
	ns, err := NewNamespace(MustAssetReader(InfraNamespaceAsset))
	if err != nil {
		panic(err)
	}
	return ns
}

func InfraMachineConfigPool() *mcfgv1.MachineConfigPool {
	mcp, err := NewMachineConfigPool(MustAssetReader(InfraMachineConfigPoolAsset))
	if err != nil {
		panic(err)
	}
	return mcp
}

func NewNamespace(manifest io.Reader) (*corev1.Namespace, error) {
	ns := corev1.Namespace{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&ns); err != nil {
		return nil, err
	}
	return &ns, nil
}

func NewMachineConfigPool(manifest io.Reader) (*mcfgv1.MachineConfigPool, error) {
	mcp := mcfgv1.MachineConfigPool{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&mcp); err != nil {
		return nil, err
	}
	return &mcp, nil
}
