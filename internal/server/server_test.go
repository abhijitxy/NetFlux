package server

import (
	"testing"
	"time"
)

type mockNetwork struct {
	server *Server
}

func (m *mockNetwork) RandomServer() *Server {
	return m.server
}

func TestNewServer(t *testing.T) {
	mockNet := &mockNetwork{}
	s := NewServer(1, mockNet)
	if s.ID != 1 {
		t.Errorf("Expected server ID 1, got %d", s.ID)
	}
	if len(s.data) != 0 {
		t.Errorf("Expected empty data map, got %d items", len(s.data))
	}
}

func TestReceiveAndGetPacket(t *testing.T) {
	mockNet := &mockNetwork{}
	s := NewServer(1, mockNet)
	packet := NewDataPacket(1, "test data")

	s.Receive(packet)

	retrievedPacket := s.GetPacket(1)
	if retrievedPacket == nil {
		t.Error("Expected to retrieve packet, got nil")
	}
	if retrievedPacket.Data != "test data" {
		t.Errorf("Expected data 'test data', got '%s'", retrievedPacket.Data)
	}
}

func TestTransmitRandomPacket(t *testing.T) {
	mockNet := &mockNetwork{}
	s := NewServer(1, mockNet)
	mockNet.server = NewServer(2, mockNet)

	packet := NewDataPacket(1, "test data")
	s.Receive(packet)

	s.transmitRandomPacket()

	// The packet should now be in the mock server
	time.Sleep(10 * time.Millisecond) // Allow some time for transmission
	retrievedPacket := mockNet.server.GetPacket(1)
	if retrievedPacket == nil {
		t.Error("Expected to retrieve packet from mock server, got nil")
	}
	if retrievedPacket.Data != "test data" {
		t.Errorf("Expected data 'test data', got '%s'", retrievedPacket.Data)
	}
}