package main_test

import (
	"strings"
	"testing"

	"github.com/jordanbecketmoore/oapirouter/pkg/oapirouter"
	"github.com/jordanbecketmoore/oapirouter/test/constants"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pb33f/libopenapi"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestOapirouter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Oapirouter Suite")
}

var _ = Describe("DocumentModelToHTTPRoute", func() {
	var (
		documentString string
		document       libopenapi.Document
		errs           []error
		httpRoute      gatewayv1.HTTPRoute
		err            error
		documentModel  *libopenapi.DocumentModel[v3.Document]
	)

	JustBeforeEach(func() {
		document, err = libopenapi.NewDocument([]byte(documentString))
		Expect(err).ToNot(HaveOccurred())

		documentModel, errs = document.BuildV3Model()
		Expect(errs).To(BeEmpty())

		httpRoute, err = oapirouter.DocumentModelToHTTPRoute(documentModel)
	})

	Context("with exact path OpenAPI spec", func() {
		BeforeEach(func() {
			documentString = strings.Join([]string{
				constants.PetstoreOpenAPI3Metadata,
				"paths:",
				constants.PetstoreOpenAPI3PathsExact,
				constants.PetstoreOpenAPI3Components,
			}, "\n")
		})

		It("should convert without error", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(httpRoute).ToNot(BeNil())
		})
	})

	Context("with path queries in OpenAPI spec", func() {
		BeforeEach(func() {
			documentString = strings.Join([]string{
				constants.PetstoreOpenAPI3Metadata,
				"paths:",
				constants.PetstoreOpenAPI3PathsQueries,
				constants.PetstoreOpenAPI3Components,
			}, "\n")
		})

		It("should convert without error", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(httpRoute).ToNot(BeNil())
		})
	})

	Context("with path parameters in OpenAPI spec", func() {
		BeforeEach(func() {
			documentString = strings.Join([]string{
				constants.PetstoreOpenAPI3Metadata,
				"paths:",
				constants.PetstoreOpenAPI3PathsRegex,
				constants.PetstoreOpenAPI3Components,
			}, "\n")
		})

		It("should convert without error", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(httpRoute).ToNot(BeNil())
		})
	})
})
