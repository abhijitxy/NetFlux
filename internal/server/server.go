package server

import (
	"sync"
	"time"
)

type DataPacket struct {
	ID   int
	Data string
}

type Server struct {
	ID      int
	data    map[int]*DataPacket
	mutex   sync.RWMutex
	network NetworkInterface
}

type NetworkInterface interface {
	RandomServer() *Server
}

func NewServer(id int, network NetworkInterface) *Server {
	return &Server{
		ID:      id,
		data:    make(map[int]*DataPacket),
		network: network,
	}
}

func (s *Server) Start() {
	go s.transmitLoop()
}

func (s *Server) transmitLoop() {
	for {
		time.Sleep(100 * time.Millisecond)
		s.transmitRandomPacket()
	}
}

func (s *Server) transmitRandomPacket() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.data) == 0 {
		return
	}

	for id, packet := range s.data {
		delete(s.data, id)
		go s.network.RandomServer().Receive(packet)
		return
	}
}

func (s *Server) Receive(packet *DataPacket) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[packet.ID] = packet
}

func (s *Server) GetPacket(id int) *DataPacket {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.data[id]
}

func NewDataPacket(id int, data string) *DataPacket {
	return &DataPacket{
		ID:   id,
		Data: data,
	}
}