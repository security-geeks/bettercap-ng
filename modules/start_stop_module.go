package modules

import (
	"fmt"
	"github.com/evilsocket/bettercap-ng/session"
)

type StartStopModule struct {
	session.SessionModule
}

func NewStartStopModule(name string, s *session.Session) StartStopModule {
	mod := StartStopModule{
		SessionModule: session.NewSessionModule(name, s),
	}

	mod.AddHandler(session.NewModuleHandler(name+" on", "",
		fmt.Sprintf("Start the %s module.", name),
		func(args []string) error {
			return mod.Start()
		}))

	mod.AddHandler(session.NewModuleHandler(name+" off", "",
		fmt.Sprintf("Stop the %s module.", name),
		func(args []string) error {
			return mod.Stop()
		}))

	return mod
}

func (mod *StartStopModule) Worker() {
	fmt.Println("wtf?!")
}

func (mod *StartStopModule) Configure() error {
	fmt.Println("wtf.configure?!")
	return nil
}

func (mod *StartStopModule) Start() error {
	if mod.Running() == true {
		return session.ErrAlreadyStarted
	} else if err := mod.Configure(); err != nil {
		return err
	}

	mod.SetRunning(true)
	go func() {
		mod.Worker()
		mod.SetRunning(false)
	}()

	return nil
}

func (mod *StartStopModule) Stop() error {
	if mod.Running() == false {
		return session.ErrAlreadyStopped
	}
	mod.SetRunning(false)
	return nil
}
