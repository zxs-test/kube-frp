package v1alpha1

// APIMetadata defines the API metadata
type APIMetadata struct {
	// Version specifies the API version
	Version string `json:"version"`
}

// AuthScope defines the authentication scope
type AuthScope string

const (
	// AuthScopeHeartBeats specifies the heartbeat authentication scope
	AuthScopeHeartBeats AuthScope = "HeartBeats"
	// AuthScopeNewWorkConns specifies the new work connections authentication scope
	AuthScopeNewWorkConns AuthScope = "NewWorkConns"
)

// AuthMethod defines the authentication method
type AuthMethod string

const (
	// AuthMethodToken specifies the token authentication method
	AuthMethodToken AuthMethod = "token"
	// AuthMethodOIDC specifies the OIDC authentication method
	AuthMethodOIDC AuthMethod = "oidc"
)

// FRPServerSpec defines the configuration for the FRP server
// +kubebuilder:object:generate=true
type FRPServerSpec struct {
	// APIMetadata specifies the API metadata
	APIMetadata `json:",inline"`

	// Auth specifies the authentication configuration
	// +optional
	Auth AuthServerConfig `json:"auth,omitempty"`

	// BindAddr specifies the address that the server binds to. By default,
	// this value is "0.0.0.0".
	// +kubebuilder:default="0.0.0.0"
	// +optional
	BindAddr string `json:"bindAddr,omitempty"`

	// BindPort specifies the port that the server listens on. By default, this
	// value is 7000.
	// +kubebuilder:default=7000
	// +optional
	BindPort int `json:"bindPort,omitempty"`

	// KCPBindPort specifies the KCP port that the server listens on. If this
	// value is 0, the server will not listen for KCP connections.
	// +optional
	KCPBindPort int `json:"kcpBindPort,omitempty"`

	// QUICBindPort specifies the QUIC port that the server listens on.
	// Set this value to 0 will disable this feature.
	// +optional
	QUICBindPort int `json:"quicBindPort,omitempty"`

	// ProxyBindAddr specifies the address that the proxy binds to. This value
	// may be the same as BindAddr.
	// +optional
	ProxyBindAddr string `json:"proxyBindAddr,omitempty"`

	// VhostHTTPPort specifies the port that the server listens for HTTP Vhost
	// requests. If this value is 0, the server will not listen for HTTP
	// requests.
	// +optional
	VhostHTTPPort int `json:"vhostHTTPPort,omitempty"`

	// VhostHTTPTimeout specifies the response header timeout for the Vhost
	// HTTP server, in seconds. By default, this value is 60.
	// +kubebuilder:default=60
	// +optional
	VhostHTTPTimeout int64 `json:"vhostHTTPTimeout,omitempty"`

	// VhostHTTPSPort specifies the port that the server listens for HTTPS
	// Vhost requests. If this value is 0, the server will not listen for HTTPS
	// requests.
	// +optional
	VhostHTTPSPort int `json:"vhostHTTPSPort,omitempty"`

	// TCPMuxHTTPConnectPort specifies the port that the server listens for TCP
	// HTTP CONNECT requests. If the value is 0, the server will not multiplex TCP
	// requests on one single port. If it's not - it will listen on this value for
	// HTTP CONNECT requests.
	// +optional
	TCPMuxHTTPConnectPort int `json:"tcpmuxHTTPConnectPort,omitempty"`

	// If TCPMuxPassthrough is true, frps won't do any update on traffic.
	// +optional
	TCPMuxPassthrough bool `json:"tcpmuxPassthrough,omitempty"`

	// SubDomainHost specifies the domain that will be attached to sub-domains
	// requested by the client when using Vhost proxying. For example, if this
	// value is set to "frps.com" and the client requested the subdomain
	// "test", the resulting URL would be "test.frps.com".
	// +optional
	SubDomainHost string `json:"subDomainHost,omitempty"`

	// Custom404Page specifies a path to a custom 404 page to display. If this
	// value is "", a default page will be displayed.
	// +optional
	Custom404Page string `json:"custom404Page,omitempty"`

	// SSHTunnelGateway specifies the SSH tunnel gateway configuration
	// +optional
	SSHTunnelGateway SSHTunnelGateway `json:"sshTunnelGateway,omitempty"`

	// WebServer specifies the web server configuration
	// +optional
	WebServer WebServerConfig `json:"webServer,omitempty"`

	// EnablePrometheus will export prometheus metrics on webserver address
	// in /metrics api.
	// +optional
	EnablePrometheus bool `json:"enablePrometheus,omitempty"`

	// Log specifies the logging configuration
	// +optional
	Log LogConfig `json:"log,omitempty"`

	// Transport specifies the transport layer configuration
	// +optional
	Transport ServerTransportConfig `json:"transport,omitempty"`

	// DetailedErrorsToClient defines whether to send the specific error (with
	// debug info) to frpc. By default, this value is true.
	// +kubebuilder:default=true
	// +optional
	DetailedErrorsToClient *bool `json:"detailedErrorsToClient,omitempty"`

	// MaxPortsPerClient specifies the maximum number of ports a single client
	// may proxy to. If this value is 0, no limit will be applied.
	// +optional
	MaxPortsPerClient int64 `json:"maxPortsPerClient,omitempty"`

	// UserConnTimeout specifies the maximum time to wait for a work
	// connection. By default, this value is 10.
	// +kubebuilder:default=10
	// +optional
	UserConnTimeout int64 `json:"userConnTimeout,omitempty"`

	// UDPPacketSize specifies the UDP packet size
	// By default, this value is 1500
	// +kubebuilder:default=1500
	// +optional
	UDPPacketSize int64 `json:"udpPacketSize,omitempty"`

	// NatHoleAnalysisDataReserveHours specifies the hours to reserve nat hole analysis data.
	// +kubebuilder:default=168
	// +optional
	NatHoleAnalysisDataReserveHours int64 `json:"natholeAnalysisDataReserveHours,omitempty"`

	// AllowPorts specifies the allowed ports for proxy
	// +optional
	AllowPorts []PortsRange `json:"allowPorts,omitempty"`

	// HTTPPlugins specifies the HTTP plugin configuration
	// +optional
	HTTPPlugins []HTTPPluginOptions `json:"httpPlugins,omitempty"`
}

// AuthServerConfig defines the authentication configuration
type AuthServerConfig struct {
	// Method specifies the authentication method
	// +kubebuilder:validation:Enum=token;oidc
	// +kubebuilder:default=token
	// +optional
	Method AuthMethod `json:"method,omitempty"`

	// AdditionalScopes specifies additional authentication scopes
	// +optional
	AdditionalScopes []AuthScope `json:"additionalScopes,omitempty"`

	// Token specifies the authentication token
	// +optional
	Token string `json:"token,omitempty"`

	// OIDC specifies the OIDC authentication configuration
	// +optional
	OIDC AuthOIDCServerConfig `json:"oidc,omitempty"`
}

// AuthOIDCServerConfig defines the OIDC authentication configuration
type AuthOIDCServerConfig struct {
	// Issuer specifies the issuer to verify OIDC tokens with. This issuer
	// will be used to load public keys to verify signature and will be compared
	// with the issuer claim in the OIDC token.
	// +optional
	Issuer string `json:"issuer,omitempty"`

	// Audience specifies the audience OIDC tokens should contain when validated.
	// If this value is empty, audience ("client ID") verification will be skipped.
	// +optional
	Audience string `json:"audience,omitempty"`

	// SkipExpiryCheck specifies whether to skip checking if the OIDC token is
	// expired.
	// +optional
	SkipExpiryCheck bool `json:"skipExpiryCheck,omitempty"`

	// SkipIssuerCheck specifies whether to skip checking if the OIDC token's
	// issuer claim matches the issuer specified in OidcIssuer.
	// +optional
	SkipIssuerCheck bool `json:"skipIssuerCheck,omitempty"`
}

// ServerTransportConfig defines the transport layer configuration
type ServerTransportConfig struct {
	// TCPMux toggles TCP stream multiplexing. This allows multiple requests
	// from a client to share a single TCP connection. By default, this value
	// is true.
	// +kubebuilder:default=true
	// +optional
	TCPMux *bool `json:"tcpMux,omitempty"`

	// TCPMuxKeepaliveInterval specifies the keep alive interval for TCP stream multiplier.
	// If TCPMux is true, heartbeat of application layer is unnecessary because it can only rely on heartbeat in TCPMux.
	// +kubebuilder:default=30
	// +optional
	TCPMuxKeepaliveInterval int64 `json:"tcpMuxKeepaliveInterval,omitempty"`

	// TCPKeepAlive specifies the interval between keep-alive probes for an active network connection between frpc and frps.
	// If negative, keep-alive probes are disabled.
	// +kubebuilder:default=7200
	// +optional
	TCPKeepAlive int64 `json:"tcpKeepAlive,omitempty"`

	// MaxPoolCount specifies the maximum pool size for each proxy. By default,
	// this value is 5.
	// +kubebuilder:default=5
	// +optional
	MaxPoolCount int64 `json:"maxPoolCount,omitempty"`

	// HeartBeatTimeout specifies the maximum time to wait for a heartbeat
	// before terminating the connection. It is not recommended to change this
	// value. By default, this value is 90. Set negative value to disable it.
	// +kubebuilder:default=90
	// +optional
	HeartbeatTimeout int64 `json:"heartbeatTimeout,omitempty"`

	// QUIC options.
	// +optional
	QUIC QUICOptions `json:"quic,omitempty"`

	// TLS specifies TLS settings for the connection from the client.
	// +optional
	TLS TLSServerConfig `json:"tls,omitempty"`
}

// QUICOptions defines the QUIC protocol options
type QUICOptions struct {
	// KeepalivePeriod specifies the keepalive period
	// +kubebuilder:default=10
	// +optional
	KeepalivePeriod int `json:"keepalivePeriod,omitempty"`

	// MaxIdleTimeout specifies the maximum idle timeout
	// +kubebuilder:default=30
	// +optional
	MaxIdleTimeout int `json:"maxIdleTimeout,omitempty"`

	// MaxIncomingStreams specifies the maximum incoming streams
	// +kubebuilder:default=100000
	// +optional
	MaxIncomingStreams int `json:"maxIncomingStreams,omitempty"`
}

// TLSServerConfig defines the TLS server configuration
type TLSServerConfig struct {
	// Force specifies whether to only accept TLS-encrypted connections.
	// +optional
	Force bool `json:"force,omitempty"`

	// TLSConfig specifies the TLS configuration
	// +optional
	TLSConfig `json:",inline"`
}

// TLSConfig defines the TLS configuration
type TLSConfig struct {
	// CertFile specifies the path of the cert file that client will load.
	// +optional
	CertFile string `json:"certFile,omitempty"`

	// KeyFile specifies the path of the secret key file that client will load.
	// +optional
	KeyFile string `json:"keyFile,omitempty"`

	// TrustedCaFile specifies the path of the trusted ca file that will load.
	// +optional
	TrustedCaFile string `json:"trustedCaFile,omitempty"`

	// ServerName specifies the custom server name of tls certificate. By
	// default, server name if same to ServerAddr.
	// +optional
	ServerName string `json:"serverName,omitempty"`
}

// SSHTunnelGateway defines the SSH tunnel gateway configuration
type SSHTunnelGateway struct {
	// BindPort specifies the port that the SSH tunnel gateway binds to
	// +optional
	BindPort int `json:"bindPort,omitempty"`

	// PrivateKeyFile specifies the path to the private key file
	// +optional
	PrivateKeyFile string `json:"privateKeyFile,omitempty"`

	// AutoGenPrivateKeyPath specifies the path to auto generate private key
	// +kubebuilder:default="./.autogen_ssh_key"
	// +optional
	AutoGenPrivateKeyPath string `json:"autoGenPrivateKeyPath,omitempty"`

	// AuthorizedKeysFile specifies the path to the authorized keys file
	// +optional
	AuthorizedKeysFile string `json:"authorizedKeysFile,omitempty"`
}

// WebServerConfig defines the web server configuration
type WebServerConfig struct {
	// This is the network address to bind on for serving the web interface and API.
	// By default, this value is "127.0.0.1".
	// +kubebuilder:default="127.0.0.1"
	// +optional
	Addr string `json:"addr,omitempty"`

	// Port specifies the port for the web server to listen on. If this
	// value is 0, the admin server will not be started.
	// +optional
	Port int `json:"port,omitempty"`

	// User specifies the username that the web server will use for login.
	// +optional
	User string `json:"user,omitempty"`

	// Password specifies the password that the admin server will use for login.
	// +optional
	Password string `json:"password,omitempty"`

	// AssetsDir specifies the local directory that the admin server will load
	// resources from. If this value is "", assets will be loaded from the
	// bundled executable using embed package.
	// +optional
	AssetsDir string `json:"assetsDir,omitempty"`

	// Enable golang pprof handlers.
	// +optional
	PprofEnable bool `json:"pprofEnable,omitempty"`

	// Enable TLS if TLSConfig is not nil.
	// +optional
	TLS *TLSConfig `json:"tls,omitempty"`
}

// LogConfig defines the logging configuration
type LogConfig struct {
	// This is destination where frp should write the logs.
	// If "console" is used, logs will be printed to stdout, otherwise,
	// logs will be written to the specified file.
	// By default, this value is "console".
	// +kubebuilder:default="console"
	// +optional
	To string `json:"to,omitempty"`

	// Level specifies the minimum log level. Valid values are "trace",
	// "debug", "info", "warn", and "error". By default, this value is "info".
	// +kubebuilder:default="info"
	// +kubebuilder:validation:Enum=trace;debug;info;warn;error
	// +optional
	Level string `json:"level,omitempty"`

	// MaxDays specifies the maximum number of days to store log information
	// before deletion.
	// +kubebuilder:default=3
	// +optional
	MaxDays int64 `json:"maxDays,omitempty"`

	// DisablePrintColor disables log colors when log.to is "console".
	// +optional
	DisablePrintColor bool `json:"disablePrintColor,omitempty"`
}

// HTTPPluginOptions defines the HTTP plugin configuration
type HTTPPluginOptions struct {
	// Name specifies the plugin name
	Name string `json:"name"`

	// Addr specifies the plugin address
	Addr string `json:"addr"`

	// Path specifies the plugin path
	Path string `json:"path"`

	// Ops specifies the plugin operations
	Ops []string `json:"ops"`

	// TLSVerify specifies whether to verify TLS
	// +optional
	TLSVerify bool `json:"tlsVerify,omitempty"`
}

// PortsRange defines a range of ports
type PortsRange struct {
	// Start specifies the start port
	Start int `json:"start"`

	// End specifies the end port
	End int `json:"end"`
}
