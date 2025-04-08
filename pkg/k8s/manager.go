package k8s

import (
	"context"
	"fmt"
	"sync"
)

// Manager handles the Kubernetes integration for the frp server
type Manager struct {
	client     Client
	serverName string
	namespace  string
	mu         sync.RWMutex
	server     *FrpServer
}

// NewManager creates a new Kubernetes manager
func NewManager(client Client, serverName, namespace string) *Manager {
	return &Manager{
		client:     client,
		serverName: serverName,
		namespace:  namespace,
	}
}

// Start begins watching the FrpServer resource
func (m *Manager) Start(ctx context.Context) error {
	// Initial fetch of the FrpServer resource
	server, err := m.client.GetFrpServer(ctx, m.serverName)
	if err != nil {
		return fmt.Errorf("failed to get FrpServer: %v", err)
	}

	m.mu.Lock()
	m.server = server
	m.mu.Unlock()

	// TODO: Implement watch using dynamic client or code-generator
	return nil
}

// GetConfig returns the current FrpServer configuration
func (m *Manager) GetConfig() (*FrpServer, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.server, nil
}

// UpdateClientConnection updates the status with a new client connection
func (m *Manager) UpdateClientConnection(ctx context.Context, conn ClientConnection) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Add the new connection to the status
	m.server.Status.ClientConnections = append(m.server.Status.ClientConnections, conn)

	// Update the status in Kubernetes
	return m.client.UpdateFrpServerStatus(ctx, m.server)
}

// RemoveClientConnection removes a client connection from the status
func (m *Manager) RemoveClientConnection(ctx context.Context, clientID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Find and remove the connection
	for i, conn := range m.server.Status.ClientConnections {
		if conn.ClientID == clientID {
			m.server.Status.ClientConnections = append(
				m.server.Status.ClientConnections[:i],
				m.server.Status.ClientConnections[i+1:]...,
			)
			break
		}
	}

	// Update the status in Kubernetes
	return m.client.UpdateFrpServerStatus(ctx, m.server)
}
