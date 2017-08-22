package proxy_test

import (
	"context"
	"testing"

	"github.com/Akagi201/spy/proxy"
)

// TestProxyFromContext -
func TestProxyFromContext(t *testing.T) {
	ctx := proxy.WithProxy(context.Background(), "a", "b")
	addrs, ok := proxy.ProxyFromContext(ctx)
	if !ok {
		t.Fail()
	}
	if len(addrs) != 2 {
		t.Fail()
	}
	if addrs[0] != "a" || addrs[1] != "b" {
		t.Fail()
	}
}
