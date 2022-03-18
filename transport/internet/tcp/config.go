package tcp

import (
	"github.com/jaskon139/xray-core/common"
	"github.com/jaskon139/xray-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
