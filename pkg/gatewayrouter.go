package oapirouter

import (
	oapi "github.com/pb33f/libopenapi"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// GatewayRouter is an object that takes openapi spec
// objects and converts them to HTTPRoutes for a given Gateway
type GatewayRouter struct {
	gatewayv1.Gateway
	routes map[oapi.Document]gatewayv1.HTTPRoute
}
