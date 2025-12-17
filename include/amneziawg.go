//go:build with_amneziawg

package include

import (
	"github.com/pulsarvpn/sing-box/adapter/endpoint"
	"github.com/pulsarvpn/sing-box/adapter/outbound"
	"github.com/pulsarvpn/sing-box/protocol/amneziawg"
)

func registerAmneziaWGOutbound(registry *outbound.Registry) {
	amneziawg.RegisterOutbound(registry)
}

func registerAmneziaWGEndpoint(registry *endpoint.Registry) {
	amneziawg.RegisterEndpoint(registry)
}
