package oapirouter

import (
	"fmt"
	"testing"

	"github.com/jordanbecketmoore/oapirouter/test/constants"
	"github.com/pb33f/libopenapi"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
	"sigs.k8s.io/yaml"
)

var TestGatewayRouter = NewGatewayRouter(constants.TestHTTPGateway)

func TestNewGatewayRouter(t *testing.T) {
	if TestGatewayRouter == nil {
		t.Errorf("Expected TestGatewayRouter to be initialized, but got nil")
	}

	if TestGatewayRouter.Gateway.Name != constants.TestHTTPGateway.Name {
		t.Errorf("Expected Gateway to be %v, but got %v", constants.TestHTTPGateway, TestGatewayRouter.Gateway)
	}
}

func TestDocumentModelToHTTPRouteExact(t *testing.T) {
	var errs []error
	documentString := constants.TestOAPI3DocStringExact
	document, _ := libopenapi.NewDocument([]byte(documentString))
	documentModel, errs := document.BuildV3Model()
	if len(errs) > 0 {
		t.Errorf("Expected no errors, but got %v", errs)
	}
	httpRoute, err := DocumentModelToHTTPRoute(documentModel)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Marhshal the HTTPRoute to YAML for better readability
	httpRouteYAML, err := yaml.Marshal(httpRoute)
	if err != nil {
		t.Errorf("Failed to marshal HTTPRoute to YAML: %v", err)
	}
	// Print the YAML
	fmt.Printf("HTTPRoute YAML:\n%s\n", string(httpRouteYAML))

	// Validate the HTTPRoute object
	scheme := runtime.NewScheme()
	err = gatewayv1.Install(scheme) // Register Gateway API types
	if err != nil {
		t.Fatalf("Failed to add Gateway API types to scheme: %v", err)
	}

	// Create a serializer for decoding and validating
	codecs := serializer.NewCodecFactory(scheme)
	decoder := codecs.UniversalDeserializer()

	// Encode the object to YAML and decode it back to validate
	_, _, err = decoder.Decode(httpRouteYAML, nil, &gatewayv1.HTTPRoute{})
	if err != nil {
		t.Errorf("HTTPRoute validation failed: %v", err)
	}
}
