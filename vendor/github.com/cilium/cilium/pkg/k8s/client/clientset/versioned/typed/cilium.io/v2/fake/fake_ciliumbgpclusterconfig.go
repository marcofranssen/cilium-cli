// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	ciliumiov2 "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/typed/cilium.io/v2"
	gentype "k8s.io/client-go/gentype"
)

// fakeCiliumBGPClusterConfigs implements CiliumBGPClusterConfigInterface
type fakeCiliumBGPClusterConfigs struct {
	*gentype.FakeClientWithList[*v2.CiliumBGPClusterConfig, *v2.CiliumBGPClusterConfigList]
	Fake *FakeCiliumV2
}

func newFakeCiliumBGPClusterConfigs(fake *FakeCiliumV2) ciliumiov2.CiliumBGPClusterConfigInterface {
	return &fakeCiliumBGPClusterConfigs{
		gentype.NewFakeClientWithList[*v2.CiliumBGPClusterConfig, *v2.CiliumBGPClusterConfigList](
			fake.Fake,
			"",
			v2.SchemeGroupVersion.WithResource("ciliumbgpclusterconfigs"),
			v2.SchemeGroupVersion.WithKind("CiliumBGPClusterConfig"),
			func() *v2.CiliumBGPClusterConfig { return &v2.CiliumBGPClusterConfig{} },
			func() *v2.CiliumBGPClusterConfigList { return &v2.CiliumBGPClusterConfigList{} },
			func(dst, src *v2.CiliumBGPClusterConfigList) { dst.ListMeta = src.ListMeta },
			func(list *v2.CiliumBGPClusterConfigList) []*v2.CiliumBGPClusterConfig {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v2.CiliumBGPClusterConfigList, items []*v2.CiliumBGPClusterConfig) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
