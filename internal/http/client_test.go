package http

import (
	"context"
	netHttp "net/http"
	"testing"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
)

func TestHttpHandlerCreateContext(t *testing.T) {
	var ctx context.Context
	server := func(w netHttp.ResponseWriter, r *netHttp.Request) {
		ctx = r.Context()
	}

	client := NewHttpClientFromHandler(netHttp.HandlerFunc(server))
	r, err := client.Get("http://localhost/foo")

	g := gomega.NewWithT(t)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(r.StatusCode).To(Equal(200))

	g.Expect(ctx).To(Equal(context.Background()))
}
