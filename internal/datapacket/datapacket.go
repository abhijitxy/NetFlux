package datapacket

import (
	"time"
)

// DataPacket represents a piece of data in the Netflux system
type DataPacket struct {
	ID        int
	Data      string
	Timestamp time.Time
}

// New creates a new DataPacket with the given ID and data
func New(id int, data string) *DataPacket {
	return &DataPacket{
		ID:        id,
		Data:      data,
		Timestamp: time.Now(),
	}
}

// UpdateTimestamp updates the timestamp of the DataPacket to the current time
func (dp *DataPacket) UpdateTimestamp() {
	dp.Timestamp = time.Now()
}

// Age returns the age of the DataPacket
func (dp *DataPacket) Age() time.Duration {
	return time.Since(dp.Timestamp)
}