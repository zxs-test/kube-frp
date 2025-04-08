package k8s

import (
	"context"
	"fmt"

	"github.com/fatedier/frp/api/v1alpha1"
	frpv1alpha1 "github.com/fatedier/frp/api/v1alpha1"
	config "github.com/fatedier/frp/pkg/config/v1"

	configtypes "github.com/fatedier/frp/pkg/config/types"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// GroupName is the group name used in this package
	GroupName = "frp.io"
	// Version is the version of the API
	Version = "v1alpha1"
)

// Client is a wrapper around the Kubernetes client
type Client struct {
	client.Client
}

// NewClient creates a new Kubernetes client
func NewClient(kubeconfig string) (*Client, error) {
	var config *rest.Config
	var err error

	if kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		config, err = rest.InClusterConfig()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get kubeconfig: %v", err)
	}

	scheme := runtime.NewScheme()
	if err := frpv1alpha1.AddToScheme(scheme); err != nil {
		return nil, fmt.Errorf("failed to add scheme: %v", err)
	}

	cl, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	return &Client{cl}, nil
}

// GetFRPServer retrieves a FRPServer CR by name and namespace
func (c *Client) GetFRPServer(ctx context.Context, name, namespace string) (*frpv1alpha1.FRPServer, error) {
	frpServer := &frpv1alpha1.FRPServer{}
	err := c.Get(ctx, client.ObjectKey{Namespace: namespace, Name: name}, frpServer)
	if err != nil {
		return nil, fmt.Errorf("failed to get FRPServer: %v", err)
	}
	return frpServer, nil
}

// LoadServerConfig loads the server configuration from a FRPServer CR
func (c *Client) LoadServerConfig(ctx context.Context, name, namespace string) (*config.ServerConfig, error) {
	frpServer, err := c.GetFRPServer(ctx, name, namespace)
	if err != nil {
		return nil, fmt.Errorf("failed to load FRPServer CR: %v", err)
	}

	// Convert FRPServerSpec to ServerConfig
	conf := &config.ServerConfig{
		Auth: config.AuthServerConfig{
			Method: config.AuthMethod(frpServer.Spec.Auth.Method),
		},
		BindAddr:              frpServer.Spec.BindAddr,
		BindPort:              frpServer.Spec.BindPort,
		KCPBindPort:           frpServer.Spec.KCPBindPort,
		QUICBindPort:          frpServer.Spec.QUICBindPort,
		ProxyBindAddr:         frpServer.Spec.ProxyBindAddr,
		VhostHTTPPort:         frpServer.Spec.VhostHTTPPort,
		VhostHTTPSPort:        frpServer.Spec.VhostHTTPSPort,
		TCPMuxHTTPConnectPort: frpServer.Spec.TCPMuxHTTPConnectPort,
		SubDomainHost:         frpServer.Spec.SubDomainHost,
		MaxPortsPerClient:     int64(frpServer.Spec.MaxPortsPerClient),
		AllowPorts:            parsePortsRange(frpServer.Spec.AllowPorts),
		EnablePrometheus:      frpServer.Spec.EnablePrometheus,
		UserConnTimeout:       int64(frpServer.Spec.UserConnTimeout),
	}

	if frpServer.Spec.Transport.TLS.Force {
		conf.Transport.TLS = config.TLSServerConfig{
			Force: true,
			TLSConfig: config.TLSConfig{
				CertFile:      frpServer.Spec.Transport.TLS.CertFile,
				KeyFile:       frpServer.Spec.Transport.TLS.KeyFile,
				TrustedCaFile: frpServer.Spec.Transport.TLS.TrustedCaFile,
			},
		}
	}

	return conf, nil
}

// parsePortsRange parses a string of ports range into a slice of PortsRange
func parsePortsRange(ports []v1alpha1.PortsRange) []configtypes.PortsRange {
	if ports == nil {
		return nil
	}
	// TODO: Implement ports range parsing
	return nil
}

func (c *Client) GetFrpServer(ctx context.Context, name string) (*FrpServer, error) {
	// TODO: Implement using dynamic client or code-generator
	return nil, fmt.Errorf("not implemented")
}

func (c *Client) UpdateFrpServerStatus(ctx context.Context, server *FrpServer) error {
	// TODO: Implement using dynamic client or code-generator
	return fmt.Errorf("not implemented")
}
