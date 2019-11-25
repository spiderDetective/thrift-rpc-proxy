package config

import (
	"net/http"
	"regexp"
	"time"
)

type Config struct {
	Proxy                Proxy
	Registry             Registry
	Listen               []Listen
	Log                  Log
	Metrics              Metrics
	UI                   UI
	Runtime              Runtime
	Tracing              Tracing
	ProfileMode          string
	ProfilePath          string
	Insecure             bool
	GlobMatchingDisabled bool
}

type CertSource struct {
	Name            string
	Type            string
	CertPath        string
	KeyPath         string
	ClientCAPath    string
	CAUpgradeCN     string
	Refresh         time.Duration
	Header          http.Header
	VaultFetchToken string
}

type Listen struct {
	Addr               string
	Proto              string
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	CertSource         CertSource
	StrictMatch        bool
	TLSMinVersion      uint16
	TLSMaxVersion      uint16
	TLSCiphers         []uint16
	ProxyProto         bool
	ProxyHeaderTimeout time.Duration
}

type UI struct {
	Listen Listen
	Color  string
	Title  string
	Access string
}

type Proxy struct {
	Strategy              string
	Matcher               string
	NoRouteStatus         int
	MaxConn               int
	ShutdownWait          time.Duration
	DialTimeout           time.Duration
	ResponseHeaderTimeout time.Duration
	KeepAliveTimeout      time.Duration
	FlushInterval         time.Duration
	GlobalFlushInterval   time.Duration
	LocalIP               string
	ClientIPHeader        string
	TLSHeader             string
	TLSHeaderValue        string
	GZIPContentTypes      *regexp.Regexp
	RequestID             string
	STSHeader             STSHeader
	AuthSchemes           map[string]AuthScheme
}

type STSHeader struct {
	MaxAge     int
	Subdomains bool
	Preload    bool
}

type Runtime struct {
	GOGC       int
	GOMAXPROCS int
}

type Circonus struct {
	APIKey        string
	APIApp        string
	APIURL        string
	CheckID       string
	BrokerID      string
	SubmissionURL string
}

type Log struct {
	AccessFormat string
	AccessTarget string
	RoutesFormat string
	Level        string
}

type Metrics struct {
	Target       string
	Prefix       string
	Names        string
	Interval     time.Duration
	Timeout      time.Duration
	Retry        time.Duration
	GraphiteAddr string
	StatsDAddr   string
	Circonus     Circonus
}

type Registry struct {
	Backend string
	Static  Static
	File    File
	Consul  Consul
	Custom  Custom
	Timeout time.Duration
	Retry   time.Duration
}

type Static struct {
	NoRouteHTML string
	Routes      string
}

type File struct {
	NoRouteHTMLPath string
	RoutesPath      string
}

type Consul struct {
	Addr                                string
	Scheme                              string
	Token                               string
	KVPath                              string
	NoRouteHTMLPath                     string
	TagPrefix                           string
	Register                            bool
	ServiceAddr                         string
	ServiceName                         string
	ServiceTags                         []string
	ServiceStatus                       []string
	CheckInterval                       time.Duration
	CheckTimeout                        time.Duration
	CheckScheme                         string
	CheckTLSSkipVerify                  bool
	CheckDeregisterCriticalServiceAfter string
	ChecksRequired                      string
	ServiceMonitors                     int
	TLS                                 ConsulTlS
}

type Custom struct {
	Host               string
	Path               string
	QueryParams        string
	Scheme             string
	CheckTLSSkipVerify bool
	PollingInterval    time.Duration
	NoRouteHTML        string
	Timeout            time.Duration
}

type Tracing struct {
	TracingEnabled bool
	CollectorType  string
	ConnectString  string
	ServiceName    string
	Topic          string
	SamplerRate    float64
	SpanHost       string
	TraceID128Bit  bool
}

type AuthScheme struct {
	Name  string
	Type  string
	Basic BasicAuth
}

type BasicAuth struct {
	Realm   string
	File    string
	Refresh time.Duration
	ModTime time.Time // the htpasswd file last modification time
}

type ConsulTlS struct {
	KeyFile            string
	CertFile           string
	CAFile             string
	CAPath             string
	InsecureSkipVerify bool
}
