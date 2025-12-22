package models

// EndpointDetails contiene información detallada del endpoint
type EndpointDetails struct {
	HostStartTime int64      `json:"hostStartTime"`
	Key           Key        `json:"key"`
	Cert          Cert       `json:"cert"`
	Chain         Chain      `json:"chain"`
	Protocols     []Protocol `json:"protocols"`
	Suites        Suites     `json:"suites"`

	// Información del servidor
	ServerSignature     string `json:"serverSignature,omitempty"`
	PrefixDelegation    bool   `json:"prefixDelegation"`
	NonPrefixDelegation bool   `json:"nonPrefixDelegation"`

	// Vulnerabilidades
	VulnBeast  bool `json:"vulnBeast"`
	Heartbleed bool `json:"heartbleed"`
	Heartbeat  bool `json:"heartbeat"`
	Poodle     bool `json:"poodle"`
	PoodleTls  int  `json:"poodleTls"`
	Freak      bool `json:"freak"`
	Logjam     bool `json:"logjam"`

	// Soporte de características
	RenegSupport             int    `json:"renegSupport"`
	SessionResumption        int    `json:"sessionResumption"`
	CompressionMethods       int    `json:"compressionMethods"`
	SupportsNpn              bool   `json:"supportsNpn"`
	NpnProtocols             string `json:"npnProtocols,omitempty"`
	SessionTickets           int    `json:"sessionTickets"`
	OcspStapling             bool   `json:"ocspStapling"`
	StaplingRevocationStatus int    `json:"staplingRevocationStatus,omitempty"`
	SniRequired              bool   `json:"sniRequired"`

	// HTTP
	HTTPStatusCode int    `json:"httpStatusCode,omitempty"`
	HTTPForwarding string `json:"httpForwarding,omitempty"`

	// RC4 y Forward Secrecy
	SupportsRc4    bool `json:"supportsRc4"`
	Rc4WithModern  bool `json:"rc4WithModern"`
	Rc4Only        bool `json:"rc4Only"`
	ForwardSecrecy int  `json:"forwardSecrecy"`

	// Otros
	OpenSslCcs         int      `json:"openSslCcs"`
	FallbackScsv       bool     `json:"fallbackScsv,omitempty"`
	HasSct             int      `json:"hasSct"`
	DhPrimes           []string `json:"dhPrimes,omitempty"`
	DhUsesKnownPrimes  int      `json:"dhUsesKnownPrimes,omitempty"`
	DhYsReuse          bool     `json:"dhYsReuse,omitempty"`
	ChaCha20Preference bool     `json:"chaCha20Preference,omitempty"`
}

// Key representa información de la clave
type Key struct {
	Size       int    `json:"size"`
	Strength   int    `json:"strength"`
	Alg        string `json:"alg"`
	DebianFlaw bool   `json:"debianFlaw"`
	Q          *int   `json:"q"`
}

// Cert representa información del certificado
type Cert struct {
	Subject              string   `json:"subject"`
	CommonNames          []string `json:"commonNames"`
	AltNames             []string `json:"altNames"`
	NotBefore            int64    `json:"notBefore"`
	NotAfter             int64    `json:"notAfter"`
	IssuerSubject        string   `json:"issuerSubject"`
	SigAlg               string   `json:"sigAlg"`
	IssuerLabel          string   `json:"issuerLabel"`
	RevocationInfo       int      `json:"revocationInfo"`
	CrlURIs              []string `json:"crlURIs"`
	OcspURIs             []string `json:"ocspURIs"`
	RevocationStatus     int      `json:"revocationStatus"`
	CrlRevocationStatus  int      `json:"crlRevocationStatus"`
	OcspRevocationStatus int      `json:"ocspRevocationStatus"`
	Sgc                  int      `json:"sgc"`
	ValidationType       string   `json:"validationType,omitempty"`
	Issues               int      `json:"issues"`
	Sct                  bool     `json:"sct"`
}

// Chain representa la cadena de certificados
type Chain struct {
	Certs  []ChainCert `json:"certs"`
	Issues int         `json:"issues"`
}

// ChainCert representa un certificado en la cadena
type ChainCert struct {
	Subject              string `json:"subject"`
	Label                string `json:"label"`
	NotBefore            int64  `json:"notBefore"`
	NotAfter             int64  `json:"notAfter"`
	IssuerSubject        string `json:"issuerSubject"`
	IssuerLabel          string `json:"issuerLabel"`
	SigAlg               string `json:"sigAlg"`
	Issues               int    `json:"issues"`
	KeyAlg               string `json:"keyAlg"`
	KeySize              int    `json:"keySize"`
	KeyStrength          int    `json:"keyStrength"`
	RevocationStatus     int    `json:"revocationStatus"`
	CrlRevocationStatus  int    `json:"crlRevocationStatus"`
	OcspRevocationStatus int    `json:"ocspRevocationStatus"`
	Raw                  string `json:"raw"`
}

// Protocol representa un protocolo soportado
type Protocol struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Version          string `json:"version"`
	V2SuitesDisabled bool   `json:"v2SuitesDisabled,omitempty"`
	Q                *int   `json:"q"`
}

// Suites contiene información de cipher suites
type Suites struct {
	List       []Suite `json:"list"`
	Preference bool    `json:"preference"`
}

// Suite representa un cipher suite
type Suite struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	CipherStrength int    `json:"cipherStrength"`
	DhStrength     int    `json:"dhStrength,omitempty"`
	DhP            int    `json:"dhP,omitempty"`
	DhG            int    `json:"dhG,omitempty"`
	DhYs           int    `json:"dhYs,omitempty"`
	EcdhBits       int    `json:"ecdhBits,omitempty"`
	EcdhStrength   int    `json:"ecdhStrength,omitempty"`
	Q              *int   `json:"q"`
}
