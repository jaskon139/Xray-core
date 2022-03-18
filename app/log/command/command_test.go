package command_test

import (
	"context"
	"testing"

	"github.com/jaskon139/xray-core/app/dispatcher"
	"github.com/jaskon139/xray-core/app/log"
	. "github.com/jaskon139/xray-core/app/log/command"
	"github.com/jaskon139/xray-core/app/proxyman"
	_ "github.com/jaskon139/xray-core/app/proxyman/inbound"
	_ "github.com/jaskon139/xray-core/app/proxyman/outbound"
	"github.com/jaskon139/xray-core/common"
	"github.com/jaskon139/xray-core/common/serial"
	"github.com/jaskon139/xray-core/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
