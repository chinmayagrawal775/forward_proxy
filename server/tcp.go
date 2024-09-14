package server

import (
	"log"
	"net"
	"os"
)

type proxyServer struct {
	server net.Listener
}

// Proxy server accessible @ http://127.0.0.1:6969/
func InitProxyServer() *proxyServer {
	listener, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Printf("Cannot start the proxy server. Error: %s", err.Error())
		os.Exit(1)
	}

	log.Println("Hello from forward proxy. Server is listening at port 6969")

	return &proxyServer{server: listener}
}

func (s *proxyServer) ShutdownServer() {
	s.server.Close()
}

func (s *proxyServer) AcceptNewConnection() net.Conn {
	conn, err := s.server.Accept()
	if err != nil {
		log.Printf("Proxy Server not accepting new connections. Error: %s", err.Error())
		return nil
	}

	return conn
}
