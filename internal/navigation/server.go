package navigation

import (
	"github.com/gorilla/websocket"
	"net"
	"net/url"
	"time"
)

const (
	propertyListener   string = "/PropertyListener"
	addListenerCommand string = "addListener"
	groundspeed        string = "/velocities/groundspeed-kt"
	verticalspeed      string = "/velocities/vertical-speed-fps"
	longitude          string = "/position/longitude-deg"
	latitude           string = "/position/latitude-deg"
	altitude           string = "/position/altitude-ft"
	altitudeagl        string = "/position/altitude-agl-ft"
	pitch              string = "/orientation/pitch-deg"
	roll               string = "/orientation/roll-deg"
	heading            string = "/orientation/heading-deg"
)

var listeners = []string{
	groundspeed,
	verticalspeed,
	longitude,
	latitude,
	altitude,
	altitudeagl,
	pitch,
	roll,
	heading,
}

type fgMessage struct {
	Path  string  `json:"path"`
	Value float64 `json:"value"`
}

type Server struct {
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
			Node    string
		}{
			Command: addListenerCommand,
			Node:    listener,
		})
	}
	return &server, nil
}

func (s *Server) GetNavigation(_ *EmptyMessage, stream Navigation_GetNavigationServer) error {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-ticker.C:
			msg, err := s.getNavigationMessage()
			if err != nil {
				return err
			}
			err = stream.Send(msg)
			if err != nil {
				return err
			}
		}
	}
}

func (s *Server) getNavigationMessage() (*NavigationMessage, error) {
	navi := NavigationMessage{}
	var msg fgMessage
	err := s.Conn.ReadJSON(&msg)
	if err != nil {
		return nil, err
	}
	switch msg.Path {
	case groundspeed:
		navi.Groundspeed = msg.Value
	case verticalspeed:
		navi.Verticalspeed = msg.Value
	case longitude:
		navi.Longitude = msg.Value
	case latitude:
		navi.Latitude = msg.Value
	case altitude:
		navi.Altitude = msg.Value
	case altitudeagl:
		navi.Altitudeagl = msg.Value
	case pitch:
		navi.Pitch = msg.Value
	case roll:
		navi.Roll = msg.Value
	case heading:
		navi.Heading = msg.Value
	}
	return &navi, nil
}
