package spec

import (
	"net/http"

	"github.com/jirawatfreedom/go-scaffold/services"
	"github.com/jirawatfreedom/go-scaffold/utils/logger"
	. "github.com/jirawatfreedom/go-scaffold/utils/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sanity", func() {
	var (
		node services.Node
		stop chan error
	)
	BeforeEach(func() {
		stop = make(chan error, 10)
		node = services.NewNode(&logger.StubLogger{})
		node.Start(&stop)
	})
	AfterEach(func() {
		node.Stop()
	})
	It("should show balances with GET /api/balance", func() {
		resp, err := http.Get("http://localhost:8080/api/balance?from=user1")
		Expect(err).toNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(ResponseBodyAsString(resp)).To(Equal("0"))
	})
	It("should transfer funds with GET /api/transfer", func() {
		resp, err := http.Get("http://localhost:8080/api/transfer?from=user1&to=user2&amount=17")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(ResponseBodyAsString(resp)).To(Equal("-17"))
		resp, err = http.Get("http://localhost:8080/api/balance?from=user2")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(ResponseBodyAsString(resp)).To(Equal("17"))
	})
})
