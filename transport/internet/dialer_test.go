package internet_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jaskon139/xray-core/common"
	"github.com/jaskon139/xray-core/common/net"
	"github.com/jaskon139/xray-core/testing/servers/tcp"
	. "github.com/jaskon139/xray-core/transport/internet"
)

func TestDialWithLocalAddr(t *testing.T) {
	server := &tcp.Server{}
	dest, err := server.Start()
	common.Must(err)
	defer server.Close()

	conn, err := DialSystem(context.Background(), net.TCPDestination(net.LocalHostIP, dest.Port), nil)
	common.Must(err)
	if r := cmp.Diff(conn.RemoteAddr().String(), "127.0.0.1:"+dest.Port.String()); r != "" {
		t.Error(r)
	}
	conn.Close()
}
