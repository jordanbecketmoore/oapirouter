package oapirouter_test

import (
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// TestHTTPGateway is a static HTTP Gateway object for testing purposes
// with a single listener on port 80.
var TestHTTPGateway = gatewayv1.Gateway{
	Spec: gatewayv1.GatewaySpec{
		GatewayClassName: "test-gateway-class",
		Listeners: []gatewayv1.Listener{
			{
				Name:     "http",
				Protocol: gatewayv1.HTTPProtocolType,
				Port:     80,
			},
		},
	},
}
