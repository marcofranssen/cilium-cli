// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	ciliumiov2 "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/typed/cilium.io/v2"
	gentype "k8s.io/client-go/gentype"
)

// fakeCiliumBGPNodeConfigs implements CiliumBGPNodeConfigInterface
type fakeCiliumBGPNodeConfigs struct {
	*gentype.FakeClientWithList[*v2.CiliumBGPNodeConfig, *v2.CiliumBGPNodeConfigList]
	Fake *FakeCiliumV2
}

func newFakeCiliumBGPNodeConfigs(fake *FakeCiliumV2) ciliumiov2.CiliumBGPNodeConfigInterface {
	return &fakeCiliumBGPNodeConfigs{
		gentype.NewFakeClientWithList[*v2.CiliumBGPNodeConfig, *v2.CiliumBGPNodeConfigList](
			fake.Fake,
			"",
			v2.SchemeGroupVersion.WithResource("ciliumbgpnodeconfigs"),
			v2.SchemeGroupVersion.WithKind("CiliumBGPNodeConfig"),
			func() *v2.CiliumBGPNodeConfig { return &v2.CiliumBGPNodeConfig{} },
			func() *v2.CiliumBGPNodeConfigList { return &v2.CiliumBGPNodeConfigList{} },
			func(dst, src *v2.CiliumBGPNodeConfigList) { dst.ListMeta = src.ListMeta },
			func(list *v2.CiliumBGPNodeConfigList) []*v2.CiliumBGPNodeConfig {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v2.CiliumBGPNodeConfigList, items []*v2.CiliumBGPNodeConfig) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
