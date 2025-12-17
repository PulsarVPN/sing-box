package option

import (
	"net/netip"

	"github.com/sagernet/sing/common/json/badoption"
)

// AmneziaWGAdvancedOptions provides advanced obfuscation options for AmneziaWG.
// These options provide more sophisticated traffic obfuscation than standard WireGuard.
type AmneziaWGAdvancedOptions struct {
	// JunkPacketCount (jc) - Number of junk packets to send before handshake
	JunkPacketCount int `json:"junk_packet_count,omitempty"`
	// JunkPacketMinSize (jmin) - Minimum size of junk packets
	JunkPacketMinSize int `json:"junk_packet_min_size,omitempty"`
	// JunkPacketMaxSize (jmax) - Maximum size of junk packets
	JunkPacketMaxSize int `json:"junk_packet_max_size,omitempty"`
	// InitPacketJunkSize (s1) - Size of junk to add to init handshake packets
	InitPacketJunkSize int `json:"init_packet_junk_size,omitempty"`
	// ResponsePacketJunkSize (s2) - Size of junk to add to response handshake packets
	ResponsePacketJunkSize int `json:"response_packet_junk_size,omitempty"`
	// CookiePacketJunkSize (s3) - Size of junk to add to cookie packets (alias: underload_packet_junk_size)
	CookiePacketJunkSize int `json:"cookie_packet_junk_size,omitempty"`
	// TransportPacketJunkSize (s4) - Size of junk to add to transport packets
	TransportPacketJunkSize int `json:"transport_packet_junk_size,omitempty"`
	// InitPacketMagicHeader (h1) - Magic header for init packets
	InitPacketMagicHeader uint32 `json:"init_packet_magic_header,omitempty"`
	// ResponsePacketMagicHeader (h2) - Magic header for response packets
	ResponsePacketMagicHeader uint32 `json:"response_packet_magic_header,omitempty"`
	// UnderloadPacketMagicHeader (h3) - Magic header for underload/cookie packets
	UnderloadPacketMagicHeader uint32 `json:"underload_packet_magic_header,omitempty"`
	// TransportPacketMagicHeader (h4) - Magic header for transport packets
	TransportPacketMagicHeader uint32 `json:"transport_packet_magic_header,omitempty"`
}

type AmneziaWGEndpointOptions struct {
	System     bool                             `json:"system,omitempty"`
	Name       string                           `json:"name,omitempty"`
	MTU        uint32                           `json:"mtu,omitempty"`
	Address    badoption.Listable[netip.Prefix] `json:"address"`
	PrivateKey string                           `json:"private_key"`
	ListenPort uint16                           `json:"listen_port,omitempty"`
	Peers      []AmneziaWGPeer                  `json:"peers,omitempty"`
	UDPTimeout badoption.Duration               `json:"udp_timeout,omitempty"`
	Workers    int                              `json:"workers,omitempty"`
	AmneziaWGAdvancedOptions
	DialerOptions
}

type AmneziaWGPeer struct {
	Address                     string                           `json:"address,omitempty"`
	Port                        uint16                           `json:"port,omitempty"`
	PublicKey                   string                           `json:"public_key,omitempty"`
	PreSharedKey                string                           `json:"pre_shared_key,omitempty"`
	AllowedIPs                  badoption.Listable[netip.Prefix] `json:"allowed_ips,omitempty"`
	PersistentKeepaliveInterval uint16                           `json:"persistent_keepalive_interval,omitempty"`
	Reserved                    []uint8                          `json:"reserved,omitempty"`
}

type LegacyAmneziaWGOutboundOptions struct {
	DialerOptions
	SystemInterface bool                             `json:"system_interface,omitempty"`
	InterfaceName   string                           `json:"interface_name,omitempty"`
	LocalAddress    badoption.Listable[netip.Prefix] `json:"local_address"`
	PrivateKey      string                           `json:"private_key"`
	Peers           []LegacyAmneziaWGPeer            `json:"peers,omitempty"`
	ServerOptions
	PeerPublicKey string  `json:"peer_public_key"`
	PreSharedKey  string  `json:"pre_shared_key,omitempty"`
	Reserved      []uint8 `json:"reserved,omitempty"`
	Workers       int     `json:"workers,omitempty"`
	MTU           uint32  `json:"mtu,omitempty"`
	AmneziaWGAdvancedOptions
}

type LegacyAmneziaWGPeer struct {
	ServerOptions
	PublicKey    string                           `json:"public_key,omitempty"`
	PreSharedKey string                           `json:"pre_shared_key,omitempty"`
	AllowedIPs   badoption.Listable[netip.Prefix] `json:"allowed_ips,omitempty"`
	Reserved     []uint8                          `json:"reserved,omitempty"`
}
