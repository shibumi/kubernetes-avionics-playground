package navigation

import (
	"github.com/gorilla/websocket"
	"net"
	"net/url"
)

const (
	propertyListener string = "/PropertyListener"
	addListenerCommand string = "addListener"
)

var listeners = []string{
	"/velocities/groundspeed-kt",
	"/velocities/vertical-speed-fps",
	"/position/longitude-deg",
	"/position/latitude-deg",
	"/position/altitude-ft",
	"/position/altitude-agl-ft",
	"/orientation/pitch-deg",
	"/orientation/roll-deg",
	"/orientation/heading-deg",
}

type Server struct{
	Conn *websocket.Conn
	UnimplementedNavigationServer
}

func NewServer(host, port string) (*Server, error) {
	server := Server{}
	addr := net.JoinHostPort(host, port)
	u := url.URL{Scheme: "ws", Host: addr, Path: propertyListener}
	var err error
	server.Conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	for _, listener := range listeners {
		err = server.Conn.WriteJSON(struct {
			Command string
			Node string
		}{
			Command: addListenerCommand,
			Node:    listener,
		})
	}
	return &server, nil
}

func (Server) GetNavigation(_ *EmptyMessage, server Navigation_GetNavigationServer) error {
	panic("implement me")
}

