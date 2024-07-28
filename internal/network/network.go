package network

import (
	"math/rand"
	"sync"
	"time"

	"github.com/abhijitxy/netflux/internal/server"
)

type Network struct {
	Servers []*server.Server
	mutex   sync.RWMutex
}

func NewNetwork(serverCount int) *Network {
	network := &Network{
		Servers: make([]*server.Server, serverCount),
	}

	for i := 0; i < serverCount; i++ {
		network.Servers[i] = server.NewServer(i, network)
	}

	return network
}

func (n *Network) Start() {
	for _, s := range n.Servers {
		go s.Start()
	}
}

func (n *Network) Insert(data string) int {
	id := rand.Intn(1000000)
	packet := server.NewDataPacket(id, data)
	randomServer := n.Servers[rand.Intn(len(n.Servers))]
	randomServer.Receive(packet)
	return id
}

func (n *Network) Retrieve(id int) (string, bool) {
	n.mutex.RLock()
	defer n.mutex.RUnlock()

	for _, s := range n.Servers {
		if packet := s.GetPacket(id); packet != nil {
			return packet.Data, true
		}
	}

	return "", false
}

func (n *Network) RandomServer() *server.Server {
	return n.Servers[rand.Intn(len(n.Servers))]
}

func init() {
	rand.Seed(time.Now().UnixNano())
}