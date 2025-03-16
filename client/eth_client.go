package client

import (
	"sync"
)

type EthClientPool struct {
	clients []string
	mu      sync.Mutex
	index   int
}

func NewEthClientPool(clients []string) *EthClientPool {
	return &EthClientPool{clients: clients}
}
func (p *EthClientPool) Len() int {
	if p.clients == nil {
		return 0
	}
	return len(p.clients)
}

func (p *EthClientPool) GetClient() (string, int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	client := p.clients[p.index]
	p.index = (p.index + 1) % len(p.clients)
	return client, p.index
}

func (p *EthClientPool) GetClients(n int) []string {
	p.mu.Lock()
	defer p.mu.Unlock()

	var selectedClients []string
	for i := 0; i < n; i++ {
		client := p.clients[(p.index+i)%len(p.clients)]
		selectedClients = append(selectedClients, client)
	}
	p.index = (p.index + n) % len(p.clients) // Update the index - round-robin method
	return selectedClients
}
