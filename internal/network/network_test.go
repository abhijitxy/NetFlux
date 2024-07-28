package network

import (
	"testing"
	"time"
)

func TestNewNetwork(t *testing.T) {
	n := NewNetwork(5)
	if len(n.Servers) != 5 {
		t.Errorf("Expected 5 servers, got %d", len(n.Servers))
	}
}

func TestInsertAndRetrieve(t *testing.T) {
	n := NewNetwork(5)
	n.Start()

	// Insert data
	testData := "test data"
	id := n.Insert(testData)

	// Wait for data to propagate
	time.Sleep(100 * time.Millisecond)

	// Retrieve data
	retrievedData, found := n.Retrieve(id)

	if !found {
		t.Errorf("Data with id %d not found", id)
	}

	if retrievedData != testData {
		t.Errorf("Expected %s, got %s", testData, retrievedData)
	}
}

func TestRetrieveNonexistent(t *testing.T) {
	n := NewNetwork(5)
	n.Start()

	_, found := n.Retrieve(999999) // Assuming this ID doesn't exist
	if found {
		t.Error("Found data for nonexistent ID")
	}
}

func TestRandomServer(t *testing.T) {
	n := NewNetwork(5)
	server := n.RandomServer()
	if server == nil {
		t.Error("RandomServer returned nil")
	}
}