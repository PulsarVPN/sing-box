//go:build with_wireguard

package include

import (
	"github.com/PulsarVPN/sing-box/adapter/endpoint"
	"github.com/PulsarVPN/sing-box/adapter/outbound"
	"github.com/PulsarVPN/sing-box/protocol/wireguard"
)

func registerWireGuardOutbound(registry *outbound.Registry) {
	wireguard.RegisterOutbound(registry)
}

func registerWireGuardEndpoint(registry *endpoint.Registry) {
	wireguard.RegisterEndpoint(registry)
}
