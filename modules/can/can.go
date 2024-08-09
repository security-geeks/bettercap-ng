package can

import (
	"net"

	"github.com/bettercap/bettercap/session"
	"go.einride.tech/can/pkg/descriptor"
	"go.einride.tech/can/pkg/socketcan"
)

type CANModule struct {
	session.SessionModule

	deviceName string
	transport  string
	dbcPath    string
	dbc        *descriptor.Database

	conn net.Conn
	recv *socketcan.Receiver
}

func NewCanModule(s *session.Session) *CANModule {
	mod := &CANModule{
		SessionModule: session.NewSessionModule("can", s),
		dbcPath:       "",
		transport:     "can",
		deviceName:    "can0",
	}

	mod.AddParam(session.NewStringParameter("can.device",
		mod.deviceName,
		"",
		"CAN-bus device."))

	mod.AddParam(session.NewStringParameter("can.transport",
		mod.transport,
		"",
		"Network type, can be 'can' for SocketCAN or 'udp'."))

	mod.AddParam(session.NewStringParameter("can.dbc_path",
		mod.dbcPath,
		"",
		"Optional path to DBC file for decoding."))

	mod.AddHandler(session.NewModuleHandler("can.recon on", "",
		"Start CAN-bus discovery.",
		func(args []string) error {
			return mod.Start()
		}))

	mod.AddHandler(session.NewModuleHandler("can.recon off", "",
		"Stop CAN-bus discovery.",
		func(args []string) error {
			return mod.Stop()
		}))

	mod.AddHandler(session.NewModuleHandler("can.clear", "",
		"Clear everything collected by the discovery module.",
		func(args []string) error {
			mod.Session.CAN.Clear()
			return nil
		}))

	mod.AddHandler(session.NewModuleHandler("can.show", "",
		"Show a list of detected CAN devices.",
		func(args []string) error {
			return mod.Show()
		}))

	return mod
}

func (mod *CANModule) Name() string {
	return "can"
}

func (mod *CANModule) Description() string {
	return "A scanner and frames injection module for CAN-bus."
}

func (mod *CANModule) Author() string {
	return "Simone Margaritelli <evilsocket@gmail.com>"
}
