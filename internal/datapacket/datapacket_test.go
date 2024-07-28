package datapacket

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	dp := New(1, "test data")
	if dp.ID != 1 {
		t.Errorf("Expected ID 1, got %d", dp.ID)
	}
	if dp.Data != "test data" {
		t.Errorf("Expected data 'test data', got '%s'", dp.Data)
	}
	if time.Since(dp.Timestamp) > time.Second {
		t.Error("Timestamp not set correctly")
	}
}

func TestUpdateTimestamp(t *testing.T) {
	dp := New(1, "test data")
	time.Sleep(10 * time.Millisecond)
	oldTimestamp := dp.Timestamp
	dp.UpdateTimestamp()
	if !dp.Timestamp.After(oldTimestamp) {
		t.Error("Timestamp not updated")
	}
}

func TestAge(t *testing.T) {
	dp := New(1, "test data")
	time.Sleep(10 * time.Millisecond)
	age := dp.Age()
	if age < 10*time.Millisecond {
		t.Errorf("Expected age to be at least 10ms, got %v", age)
	}
}