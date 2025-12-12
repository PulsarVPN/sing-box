//go:build with_wireguard

package include

import (
	"github.com/pulsarvpn/sing-box/adapter/endpoint"
	"github.com/pulsarvpn/sing-box/adapter/outbound"
	"github.com/pulsarvpn/sing-box/protocol/wireguard"
)

func registerWireGuardOutbound(registry *outbound.Registry) {
	wireguard.RegisterOutbound(registry)
}

func registerWireGuardEndpoint(registry *endpoint.Registry) {
	wireguard.RegisterEndpoint(registry)
}
