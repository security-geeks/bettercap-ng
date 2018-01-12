package modules

import (
	// "context"
	"fmt"
	"net/http"
	// "time"

	"github.com/evilsocket/bettercap-ng/log"
	"github.com/evilsocket/bettercap-ng/session"
)

type HttpServer struct {
	StartStopModule
	server *http.Server
}

func NewHttpServer(s *session.Session) *HttpServer {
	httpd := &HttpServer{
		StartStopModule: NewStartStopModule("http.server", s),
		server:          &http.Server{},
	}

	httpd.AddParam(session.NewStringParameter("http.server.path",
		".",
		"",
		"Server folder."))

	httpd.AddParam(session.NewStringParameter("http.server.address",
		session.ParamIfaceAddress,
		`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`,
		"Address to bind the http server to."))

	httpd.AddParam(session.NewIntParameter("http.server.port",
		"80",
		"Port to bind the http server to."))

	return httpd
}

func (httpd *HttpServer) Name() string {
	return "http.server"
}

func (httpd *HttpServer) Description() string {
	return "A simple HTTP server, to be used to serve files and scripts accross the network."
}

func (httpd *HttpServer) Author() string {
	return "Simone Margaritelli <evilsocket@protonmail.com>"
}

func (httpd *HttpServer) Configure() error {
	httpd.StartStopModule.Configure()

	var err error
	var path string
	var address string
	var port int

	if err, path = httpd.StringParam("http.server.path"); err != nil {
		return err
	}

	http.Handle("/", http.FileServer(http.Dir(path)))

	if err, address = httpd.StringParam("http.server.addr"); err != nil {
		return err
	}

	if err, port = httpd.IntParam("http.server.port"); err != nil {
		return err
	}

	httpd.server.Addr = fmt.Sprintf("%s:%d", address, port)

	return nil
}

func (httpd *HttpServer) Worker() {
	httpd.StartStopModule.Worker()

	log.Info("httpd server starting on http://%s", httpd.server.Addr)
	err := httpd.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

/*
func (httpd *HttpServer) Stop() error {
	if err := httpd.StartStopModule.Stop(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return httpd.server.Shutdown(ctx)
}
*/
