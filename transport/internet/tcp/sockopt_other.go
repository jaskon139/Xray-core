//go:build !linux && !freebsd
// +build !linux,!freebsd

package tcp

import (
	"github.com/jaskon139/xray-core/common/net"
	"github.com/jaskon139/xray-core/transport/internet/stat"
)

func GetOriginalDestination(conn stat.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
