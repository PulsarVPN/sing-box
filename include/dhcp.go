//go:build with_dhcp

package include

import (
	"github.com/pulsarvpn/sing-box/dns"
	"github.com/pulsarvpn/sing-box/dns/transport/dhcp"
)

func registerDHCPTransport(registry *dns.TransportRegistry) {
	dhcp.RegisterTransport(registry)
}
