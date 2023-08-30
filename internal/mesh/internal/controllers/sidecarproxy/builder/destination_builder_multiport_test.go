// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package builder

import (
	"fmt"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v1alpha1"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hashicorp/consul/internal/catalog"
	"github.com/hashicorp/consul/internal/mesh/internal/types/intermediate"
	"github.com/hashicorp/consul/internal/resource/resourcetest"
	pbmesh "github.com/hashicorp/consul/proto-public/pbmesh/v1alpha1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

func TestBuildMultiportImplicitDestinations(t *testing.T) {
	const (
		apiApp      = "api-app"
		trustDomain = "foo.consul"
		datacenter  = "dc1"
	)
	multiportEndpointsData := &pbcatalog.ServiceEndpoints{
		Endpoints: []*pbcatalog.Endpoint{
			{
				Addresses: []*pbcatalog.WorkloadAddress{
					{Host: "10.0.0.1"},
				},
				Ports: map[string]*pbcatalog.WorkloadPort{
					"admin-port": {Port: 8080, Protocol: pbcatalog.Protocol_PROTOCOL_TCP},
					"api-port":   {Port: 9090, Protocol: pbcatalog.Protocol_PROTOCOL_TCP},
					"mesh":       {Port: 20000, Protocol: pbcatalog.Protocol_PROTOCOL_MESH},
				},
			},
		},
	}
	apiAppEndpoints := resourcetest.Resource(catalog.ServiceEndpointsType, apiApp).
		WithOwner(resourcetest.Resource(catalog.ServiceType, apiApp).ID()).
		WithData(t, multiportEndpointsData).Build()

	apiAppIdentity := &pbresource.Reference{
		Name:    fmt.Sprintf("%s-identity", apiApp),
		Tenancy: apiAppEndpoints.Id.Tenancy,
	}

	proxyCfg := &pbmesh.ProxyConfiguration{
		DynamicConfig: &pbmesh.DynamicConfig{
			Mode: pbmesh.ProxyMode_PROXY_MODE_TRANSPARENT,
			TransparentProxy: &pbmesh.TransparentProxy{
				OutboundListenerPort: 15001,
			},
		},
	}

	destination1 := &intermediate.Destination{
		ServiceEndpoints: &intermediate.ServiceEndpoints{
			Resource:  apiAppEndpoints,
			Endpoints: multiportEndpointsData,
		},
		Identities: []*pbresource.Reference{apiAppIdentity},
		VirtualIPs: []string{"1.1.1.1"},
	}

	cases := map[string]struct {
		destinations []*intermediate.Destination
	}{
		"destination/l4-multiport-single-implicit-destination-tproxy": {
			destinations: []*intermediate.Destination{destination1},
		},
	}

	for name, c := range cases {
		proxyTmpl := New(testProxyStateTemplateID(), testIdentityRef(), trustDomain, datacenter, proxyCfg).
			BuildDestinations(c.destinations).
			Build()

		actual := protoToJSON(t, proxyTmpl)
		expected := goldenValue(t, name, actual, *update)

		require.JSONEq(t, expected, actual)
	}
}
