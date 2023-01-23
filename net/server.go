package net

import (
	"fmt"
	"log"
	"net"
)

// Server module
type Server struct {
	// the name of server
	Name string
	// IP Version
	IPVersion string
	// IP address
	IP string
	// port
	Port int
}

func (s *Server) Start() {
	log.Printf("[BubbleServer] [Starting]")
	go func() {
		// get tcp addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Fatalf("[BubbleServer] - start error: get tcp addr error\n")
			return
		}
		// try to listen
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Fatalf("[BubbleServer] - start error: listen to addr error\n")
			return
		}
		log.Println("[BubbleServer] - start success")
		// wait for clients
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				log.Println("[BubbleServer] - start error: accept tcp connection error")
				continue
			}
			// work
			go func() {
				for {
					buff := make([]byte, 512)
					cnt, err := conn.Read(buff)
					if err != nil {
						log.Println("[BubbleServer] - read error")
						continue
					}

					// recall
					if _, err := conn.Write(buff[:cnt]); err != nil {
						log.Println("[BubbleServer] - recall error")
						continue
					}
				}
			}()

		}
	}()

}

func (s *Server) Stop() {
}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func NewServer(name string) *Server {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}
