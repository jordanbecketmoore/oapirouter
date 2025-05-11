package oapirouter

import (
	"fmt"
	"strings"

	oapi "github.com/pb33f/libopenapi"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// GatewayRouter is an object that takes openapi spec
// objects and converts them to HTTPRoutes for a given Gateway
type GatewayRouter struct {
	gatewayv1.Gateway
	routes map[oapi.Document]gatewayv1.HTTPRoute
}

// NewGatewayRouter creates a new GatewayRouter object
// with the provided Gateway and initializes the routes map.
func NewGatewayRouter(gateway gatewayv1.Gateway) *GatewayRouter {
	return &GatewayRouter{
		Gateway: gateway,
		routes:  make(map[oapi.Document]gatewayv1.HTTPRoute),
	}
}

// DocumentModelToHTTPRoute converts an OpenAPI DocumentModel to a Gateway API HTTPRoute.
// This function only supports OpenAPI version 3.
func DocumentModelToHTTPRoute(model *oapi.DocumentModel[v3high.Document]) (gatewayv1.HTTPRoute, error) {
	// Check if the DocumentModel is for OpenAPI version 3

	// Initialize an HTTPRoute object
	var httpRoute gatewayv1.HTTPRoute

	httpRoute.Name = ToKubernetesResourceName(model.Model.Info.Title)

	// Extract paths from the DocumentModel
	paths := model.Model.Paths
	if paths == nil || paths.PathItems.IsZero() {
		return httpRoute, fmt.Errorf("no paths found in the OpenAPI document model")
	}

	// Iterate over paths and create HTTPRoute rules
	for path, pathItem := range paths.PathItems.FromNewest() {
		for method, operation := range pathItem.GetOperations().FromNewest() {

			// Create matchType
			var matchType gatewayv1.PathMatchType
			// Set method for exact match
			switch {
			case HasNoParameters(operation):
				matchType = gatewayv1.PathMatchExact
			case HasPathParameters(operation):
				matchType = gatewayv1.PathMatchRegularExpression
			case HasQueryParameters(operation):
				matchType = gatewayv1.PathMatchPathPrefix
			}

			// Create method
			method := gatewayv1.HTTPMethod(strings.ToUpper(method))

			// If path contains parameters, convert to regex
			regexpPath := ToRegularExpressionPath(path)
			// Create HTTPRouteMatch
			match := gatewayv1.HTTPRouteMatch{
				Path: &gatewayv1.HTTPPathMatch{
					Type:  &matchType,
					Value: &regexpPath,
				},
				Method: &method,
			}

			// Create HTTPRouteRule
			rule := gatewayv1.HTTPRouteRule{
				Matches: []gatewayv1.HTTPRouteMatch{match},
			}

			// Add the rule to the HTTPRoute
			httpRoute.Spec.Rules = append(httpRoute.Spec.Rules, rule)
		}
	}

	// Return the constructed HTTPRoute
	return httpRoute, nil
}
