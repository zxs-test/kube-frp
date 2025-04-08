package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FrpServer is the Schema for the frpservers API
type FrpServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FrpServerSpec   `json:"spec,omitempty"`
	Status FrpServerStatus `json:"status,omitempty"`
}

// FrpServerSpec defines the desired state of FrpServer
type FrpServerSpec struct {
	// ServiceAddress is the address where the frp server will be exposed
	ServiceAddress string `json:"serviceAddress"`
	// ServicePort is the port where the frp server will listen (default: 7000)
	ServicePort int `json:"servicePort"`
	// AuthToken is used for authentication between client and server
	AuthToken string `json:"authToken"`
	// AllowedPorts is a list of ports that are allowed to be exposed
	AllowedPorts []int `json:"allowedPorts"`
}

// FrpServerStatus defines the observed state of FrpServer
type FrpServerStatus struct {
	// ClientConnections contains information about connected clients
	ClientConnections []ClientConnection `json:"clientConnections,omitempty"`
}

// ClientConnection represents a connected client
type ClientConnection struct {
	// ClientID is the unique identifier of the client
	ClientID string `json:"clientID"`
	// LocalPort is the port exposed on the service address
	LocalPort int `json:"localPort"`
	// RemotePort is the port on the client side
	RemotePort int `json:"remotePort"`
	// Protocol is the protocol being used (tcp/udp)
	Protocol string `json:"protocol"`
	// ConnectedAt is the timestamp when the client connected
	ConnectedAt metav1.Time `json:"connectedAt"`
}

// FrpServerList contains a list of FrpServer
type FrpServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrpServer `json:"items"`
}
