package oapirouter_test

import (
	"oapirouter"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var TestGatewayRouter = oapirouter.NewGatewayRouter(TestHTTPGateway)

var _ = Describe("GatewayRouter", func() {
	When("a valid OpenAPI document string is provided", func() {
		doc := TestOAPI3DocString
		It("should generate the correct HTTPRoute", func() {
			gatewayRouter.RouteDocument(doc)
		})
	})
})
