package udp

import (
	"github.com/jaskon139/xray-core/common"
	"github.com/jaskon139/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
