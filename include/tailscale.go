//go:build with_tailscale

package include

import (
	"github.com/PulsarVPN/sing-box/adapter/endpoint"
	"github.com/PulsarVPN/sing-box/adapter/service"
	"github.com/PulsarVPN/sing-box/dns"
	"github.com/PulsarVPN/sing-box/protocol/tailscale"
	"github.com/PulsarVPN/sing-box/service/derp"
)

func registerTailscaleEndpoint(registry *endpoint.Registry) {
	tailscale.RegisterEndpoint(registry)
}

func registerTailscaleTransport(registry *dns.TransportRegistry) {
	tailscale.RegistryTransport(registry)
}

func registerDERPService(registry *service.Registry) {
	derp.Register(registry)
}
