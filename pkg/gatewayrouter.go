package oapirouter

import (
	"fmt"

	"github.com/blang/semver"
	oapi "github.com/pb33f/libopenapi"
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

// DocumentToHTTPRoute converts an OpenAPI document to a Gateway API HTTPRoute
// and returns it along with any errors encountered during the conversion.
func DocumentToHTTPRoute(doc oapi.Document) (gatewayv1.HTTPRoute, []error) {
	// Define a variable to hold any errors
	var errs []error
	// Create documentModel to hold the OpenAPI documentModel
	var documentModel oapi.DocumentModel
	// Check which version of the OpenAPI spec is being used
	version, err := semver.Parse(doc.GetVersion())
	if err != nil {
		err := fmt.Errorf("invalid OpenAPI version: %s", doc.GetVersion())
		return gatewayv1.HTTPRoute{}, append(errs, err)
	}

	switch version.Major {
	case 3:
		// OpenAPI 3.x.x to DocumentModel
		documentModel, errs = doc.BuildV3Model()
		if len(errs) > 0 {
			return gatewayv1.HTTPRoute{}, errs
		}
	case 2:
		// OpenAPI 2.x.x to DocumentModel
		documentModel, errs = doc.BuildV2Model()
		if len(errs) > 0 {
			return gatewayv1.HTTPRoute{}, errs
		}
	default:
		// Handle unsupported versions
		err := fmt.Errorf("unsupported OpenAPI version: %s", doc.GetVersion())
		return gatewayv1.HTTPRoute{}, append(errs, err)
	}

	// Placeholder return
	return gatewayv1.HTTPRoute{}, errs
}

// RouteDocument takes an OpenAPI document and adds it to the GatewayRouter
// as an HTTPRoute.
func (g *GatewayRouter) RouteDocument(doc oapi.Document) error {
	// Convert the OpenAPI document to a Gateway API HTTPRoute
	route, err := DocumentToHTTPRoute(doc)
	if err != nil {
		return err
	}

	// Add the route to the GatewayRouter's routes map
	g.routes[doc] = route

	return nil
}
